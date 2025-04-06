package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"

	// Timetable API
	"middleware/example/internal/controllers/timetable"
	timetableSvc "middleware/example/internal/services/timetable"

	// Helpers (DB, NATS)
	"middleware/example/internal/helpers"

	// Consumer
	"middleware/example/internal/consumer"
)

func runAPI() {
	// Définir le chemin ICS (avec valeur par défaut)
	icsPath := os.Getenv("ICS_FILE_PATH")
	if icsPath == "" {
		icsPath = "data/timetable.ics"
	}
	timetableService := timetableSvc.NewService(icsPath)

	// Ouvrir la base de données et initialiser le schéma
	db, err := helpers.OpenDB("file:collections.db")
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la base de données : %v", err)
	}
	defer db.Close()

	if err := helpers.InitSchema(db); err != nil {
		log.Fatalf("Erreur lors de l'initialisation du schéma : %v", err)
	}

	// Configurer les routes HTTP
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Route pour récupérer les événements (Timetable API)
	r.Get("/events", timetable.GetEvents(timetableService))

	// Lancer le serveur HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("API server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func runConsumer() {
	// Démarrer le consumer NATS (qui va récupérer et traiter les événements)
	if err := consumer.StartConsumer(); err != nil {
		log.Fatalf("Erreur lors du démarrage du consumer : %v", err)
	}
}

func main() {
	// Initialiser la connexion NATS
	helpers.InitNats()

	// Exécuter le consumer et l'API en parallèle
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		runConsumer()
	}()

	go func() {
		defer wg.Done()
		runAPI()
	}()

	wg.Wait()
}
