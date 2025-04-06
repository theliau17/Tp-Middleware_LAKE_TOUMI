package consumer

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

// --------------------------------------------------------------------------------
// 1. Templates embarqués
// Nous supposons que vos templates se trouvent dans le dossier "internal/consumer/templates".
// La directive ci-dessous embarque tous les fichiers de ce dossier.
// --------------------------------------------------------------------------------

//go:embed templates/*
var embeddedTemplates embed.FS

// MailMatter représente les informations extraites du front-matter des templates de mail.
type MailMatter struct {
	Subject string `yaml:"subject"`
}

// GetStringFromEmbeddedTemplate parse un template embarqué et retourne le contenu formaté,
// le front-matter (MailMatter) et une éventuelle erreur.
func GetStringFromEmbeddedTemplate(templatePath string, body interface{}) (content string, matter MailMatter, err error) {
	// Ajout automatique du préfixe "templates/" car les fichiers sont dans ce dossier
	fullPath := "templates/" + templatePath

	tmpl, err := template.ParseFS(embeddedTemplates, fullPath)
	if err != nil {
		return
	}

	var tpl bytes.Buffer
	if err = tmpl.Execute(&tpl, body); err != nil {
		return
	}

	// Extraction du front matter et du contenu du template
	var fm MailMatter
	mailContent, err := frontmatter.Parse(strings.NewReader(tpl.String()), &fm)
	if err == nil {
		content = string(mailContent)
		matter = fm
	}
	return
}

// --------------------------------------------------------------------------------
// 2. Fonction d'envoi d'e-mail via l'API Mail (GCC)
// --------------------------------------------------------------------------------

func SendMail(to string, subject string, textContent string, htmlContent string) error {
	// Remplacez par votre token d'authentification réel validé pour l'adresse "theliau.lake@etu.uca.fr"
	token := "yPxaSUEhxOQEWQGdemVPBGjYZyFfByFjQyHmxckB"
	// URL d'API d'après la spécification OpenAPI
	mailAPIURL := "https://mail-api.edu.forestier.re"

	// Construction du payload (ajustez les clés selon la spec de l'API)
	payload := map[string]interface{}{
		"from":      "theliau.lake@etu.uca.fr", // Doit correspondre au token
		"recipient": to,                        // Si l'API attend "recipient" au lieu de "to"
		"subject":   subject,
		"content":   htmlContent, // On envoie le contenu HTML (vous pouvez combiner texte et HTML si nécessaire)
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", mailAPIURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("échec de l'envoi du mail, statut: %d, réponse: %s", resp.StatusCode, string(bodyBytes))
	}
	return nil
}

// --------------------------------------------------------------------------------
// 3. Consumer Alerter
// Ce consumer s'abonne sur "events.modified" et envoie un mail d'alerte pour chaque message reçu.
// --------------------------------------------------------------------------------

func StartAlerterConsumer() error {
	js, err := helpers.NatsConn.JetStream()
	if err != nil {
		return fmt.Errorf("erreur JetStream: %v", err)
	}

	_, err = js.Subscribe("events.modified", func(m *nats.Msg) {
		logrus.Infof("Alerter Consumer - message reçu sur %s: %s", m.Subject, string(m.Data))

		var alertData map[string]interface{}
		if err := json.Unmarshal(m.Data, &alertData); err != nil {
			logrus.Errorf("Erreur de désérialisation du message d'alerte: %v", err)
			m.Nak()
			return
		}

		eventUID, _ := alertData["uid"].(string)
		eventName, _ := alertData["eventName"].(string)
		if eventName == "" {
			eventName = "Événement inconnu"
		}

		// Simulation d'une alerte (vous pouvez récupérer les alertes via l'API Config)
		alert := models.Alert{
			ID:         "alert-1",
			Recipients: []string{"liza.toumi@uca.fr"},
			ResourceID: "resource-1",
			Condition:  "modification",
			CreatedAt:  time.Now().UTC(),
		}

		bodyData := map[string]interface{}{
			"EventName": eventName,
			"UID":       eventUID,
		}

		textContent, mailMatter, err := GetStringFromEmbeddedTemplate("mail.txt", bodyData)
		if err != nil {
			logrus.Errorf("Erreur lors du parsing du template texte: %v", err)
			m.Nak()
			return
		}

		htmlContent, _, err := GetStringFromEmbeddedTemplate("mail.html", bodyData)
		if err != nil {
			logrus.Errorf("Erreur lors du parsing du template HTML: %v", err)
			m.Nak()
			return
		}

		for _, recipient := range alert.Recipients {
			err := SendMail(recipient, mailMatter.Subject, textContent, htmlContent)
			if err != nil {
				logrus.Errorf("Erreur lors de l'envoi du mail à %s: %v", recipient, err)
			} else {
				logrus.Infof("Mail envoyé avec succès à %s", recipient)
			}
		}

		m.Ack()
	},
		nats.Durable("alerter_consumer"),
		nats.ManualAck(),
		nats.AckExplicit(),
	)
	if err != nil {
		return fmt.Errorf("erreur lors de l'abonnement au consumer alerter: %v", err)
	}

	logrus.Info("Consumer Alerter démarré sur 'events.modified' (durable: alerter_consumer)")
	<-context.Background().Done()
	return nil
}
