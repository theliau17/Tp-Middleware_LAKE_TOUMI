package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"

	// Timetable
	"middleware/example/internal/controllers/timetable"
	timetableSvc "middleware/example/internal/services/timetable"

	// Config (Resources)
	configCtrlResource "middleware/example/internal/controllers/config"
	configRepoResource "middleware/example/internal/repositories/resource"
	configSvcResource "middleware/example/internal/services/resource"

	// Helpers
	"middleware/example/internal/helpers"
)

func main() {
	// PARTIE TIMETABLE
	icsPath := os.Getenv("ICS_FILE_PATH")
	if icsPath == "" {
		icsPath = "data/timetable.ics"
	}
	timetableService := timetableSvc.NewService(icsPath)

	// PARTIE CONFIG (DB + CRUD)
	db, err := helpers.OpenDB("file:collections.db")
	if err != nil {
		log.Fatalf("Error opening DB: %v", err)
	}
	defer db.Close()

	if err := helpers.InitSchema(db); err != nil {
		log.Fatalf("Error initializing schema: %v", err)
	}

	resourceRepo := configRepoResource.NewRepository(db)
	resourceService := configSvcResource.NewService(resourceRepo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Routes Timetable
	r.Get("/events", timetable.GetEvents(timetableService))

	// Routes Config (resources)
	r.Route("/resources", func(r chi.Router) {
		r.Get("/", configCtrlResource.GetResources(resourceService))
		r.Get("/{id}", configCtrlResource.GetResource(resourceService))
		r.Post("/", configCtrlResource.CreateResource(resourceService))
		r.Put("/{id}", configCtrlResource.UpdateResource(resourceService))
		r.Delete("/{id}", configCtrlResource.DeleteResource(resourceService))
	})

	// Lancement du serveur
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("Server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
