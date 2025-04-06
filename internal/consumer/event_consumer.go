package consumer

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	eventRepo "middleware/example/internal/repositories/event"
)

// StartConsumer démarre un consumer JetStream qui s'abonne sur le sujet "events.new"
// avec un consumer durable et utilise le dépôt et le comparateur d'événements pour traiter les messages.
func StartConsumer() error {
	// Obtenir le contexte JetStream depuis la connexion NATS
	js, err := helpers.NatsConn.JetStream()
	if err != nil {
		return fmt.Errorf("erreur JetStream: %v", err)
	}

	// Assurer l'existence du stream "EVENTS"
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "EVENTS",
		Subjects: []string{"events.*"},
	})
	if err != nil {
		logrus.Warnf("Stream EVENTS : %v", err)
	}

	// Ouvrir la base de données et créer une instance du dépôt d'événements
	db, err := helpers.OpenDB("file:collections.db")
	if err != nil {
		return fmt.Errorf("Erreur d'ouverture de la base de données: %v", err)
	}
	// Note : ne fermez pas db ici, la connexion doit rester active tant que le consumer tourne
	repo := eventRepo.NewRepository(db)

	// S'abonner sur le sujet "events.new" avec options de consumer durable et ACK manuel
	_, err = js.Subscribe("events.new", func(m *nats.Msg) {
		logrus.Debugf("Reçu message sur %s: %s", m.Subject, string(m.Data))

		// Désérialiser le message JSON dans une instance de models.Event
		var evt models.Event
		if err := json.Unmarshal(m.Data, &evt); err != nil {
			logrus.Errorf("Erreur lors de la désérialisation de l'événement: %v", err)
			m.Nak() // Indiquer que le message n'a pas pu être traité
			return
		}

		// Récupérer l'événement existant dans la base de données via son UID
		existingEvt, err := repo.GetByUID(evt.UID)
		if err != nil {
			logrus.Errorf("Erreur lors de la récupération de l'événement existant: %v", err)
			m.Nak()
			return
		}

		// Comparer l'événement existant avec le nouvel événement reçu
		if existingEvt != nil && helpers.HasChanged(*existingEvt, evt) {
			// L'événement existe et a changé : mettre à jour l'événement dans la BDD
			if err := repo.Upsert(evt); err != nil {
				logrus.Errorf("Erreur lors de la mise à jour de l'événement: %v", err)
				m.Nak()
				return
			}
			// Publier une alerte sur "events.modified"
			alertPayload, err := json.Marshal(map[string]interface{}{
				"uid":     evt.UID,
				"message": "L'événement a été modifié",
			})
			if err != nil {
				logrus.Errorf("Erreur lors de la sérialisation de l'alerte: %v", err)
			} else {
				pubAckFuture, err := js.PublishAsync("events.modified", alertPayload)
				if err != nil {
					logrus.Errorf("Erreur lors de la publication de l'alerte: %v", err)
				} else {
					select {
					case <-pubAckFuture.Ok():
						logrus.Infof("Alerte publiée pour l'événement %s", evt.UID)
					case <-pubAckFuture.Err():
						logrus.Errorf("Erreur lors de la confirmation de l'alerte pour l'événement %s", evt.UID)
					}
				}
			}
		} else if existingEvt == nil {
			// Si l'événement n'existe pas, l'insérer dans la base de données
			if err := repo.Upsert(evt); err != nil {
				logrus.Errorf("Erreur lors de l'insertion de l'événement: %v", err)
				m.Nak()
				return
			}
		}

		// Log de l'événement traité
		logrus.Infof("Event traité : UID=%s, Summary=%s, Start=%s, End=%s, Location=%s",
			evt.UID, evt.Summary, evt.StartTime, evt.EndTime, evt.Location)

		// Confirmer la réception du message
		m.Ack()
	},
		nats.Durable("timetable_consumer"),
		nats.ManualAck(),
		nats.AckExplicit(),
	)
	if err != nil {
		return fmt.Errorf("erreur lors de l'abonnement : %v", err)
	}

	logrus.Infof("Consumer démarré sur 'events.new' (durable: timetable_consumer)")
	// Bloquer indéfiniment pour maintenir le consumer actif
	select {}
}
