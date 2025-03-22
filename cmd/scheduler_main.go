package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zhashkevych/scheduler"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
	configRepoResource "middleware/example/internal/repositories/resource"
	configSvcResource "middleware/example/internal/services/resource"
)

var resourceService configSvcResource.Service

func initDBAndService() {
	db, err := helpers.OpenDB("file:collections.db")
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	if err := helpers.InitSchema(db); err != nil {
		log.Fatalf("Error initializing schema: %v", err)
	}
	repo := configRepoResource.NewRepository(db)
	resourceService = configSvcResource.NewService(repo)
}

func retrieveAndProcess(context.Context) {
	resources, err := resourceService.GetResources()
	if err != nil {
		log.Printf("[Scheduler] Erreur GetResources: %v", err)
		return
	}
	for _, res := range resources {
		parseOneResource(res)
	}
}

func parseOneResource(res models.Resource) {
	// Si file_path est vide, on lui assigne une URL par défaut
	if res.FilePath == "" {
		res.FilePath = "https://edt.uca.fr/jsp/custom/modules/plannings/anonymous_cal.jsp?resources=13295,13345&projectId=2&calType=ical&nbWeeks=1&displayConfigId=128"
	}

	// Utiliser http.Get pour récupérer le contenu ICS depuis l'URL
	resp, err := http.Get(res.FilePath)
	if err != nil {
		log.Printf("[Scheduler] Erreur http GET resource %s: %v", res.ID, err)
		return
	}
	defer resp.Body.Close()

	events, err := helpers.ParseICS(resp.Body)
	if err != nil {
		log.Printf("[Scheduler] Erreur parsing ICS resource %s: %v", res.ID, err)
		return
	}

	log.Printf("[Scheduler] Resource %s => %d événements", res.ID, len(events))
	// Afficher le détail
	for i, evt := range events {
		log.Printf("Event #%d => UID: %s, Summary: %s, Start: %s, End: %s",
			i, evt.UID, evt.Summary, evt.StartTime, evt.EndTime)
	}
}

func main() {
	initDBAndService()

	sc := scheduler.NewScheduler()
	ctx := context.Background()

	// Appelle retrieveAndProcess toutes les 2 minutes
	sc.Add(ctx, retrieveAndProcess, 2*time.Minute)

	fmt.Println("Scheduler démarré. Ctrl+C pour arrêter.")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	sc.Stop()
	fmt.Println("Scheduler arrêté.")
}
