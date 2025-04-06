package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/mattn/go-sqlite3"

	// Timetable API
	timetableCtrl "middleware/example/internal/controllers/timetable"
	timetableSvc "middleware/example/internal/services/timetable"

	// Config API (Resources)
	configCtrl "middleware/example/internal/controllers/config"
	configRepo "middleware/example/internal/repositories/resource"
	configSvc "middleware/example/internal/services/resource"

	// Helpers (DB, NATS)
	"middleware/example/internal/helpers"

	// Consumer d'événements
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

	// Créer le routeur principal
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// 1. Routes pour l'emploi du temps sous /timetable_api/events
	r.Route("/timetable_api", func(r chi.Router) {
		r.Get("/events", timetableCtrl.GetEvents(timetableService))
	})

	// 2. Routes pour l'API de configuration sous /config_api
	resourceRepo := configRepo.NewRepository(db)
	resourceService := configSvc.NewService(resourceRepo)

	r.Route("/config_api", func(r chi.Router) {
		// Routes pour les ressources
		r.Route("/resources", func(r chi.Router) {
			r.Get("/", configCtrl.GetResources(resourceService))
			r.Get("/{id}", configCtrl.GetResource(resourceService))
			r.Post("/", configCtrl.CreateResource(resourceService))
			r.Put("/{id}", configCtrl.UpdateResource(resourceService))
			r.Delete("/{id}", configCtrl.DeleteResource(resourceService))
		})
		// Routes pour les alertes (placeholders)
		r.Route("/alerts", func(r chi.Router) {
			// Ici, nous utilisons des fonctions placeholder.
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte("[]"))
			})
			r.Post("/", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"message":"Alert created (placeholder)"}`))
			})
		})
	})

	// 3. Servir les fichiers statiques du front
	// On suppose que le dossier de build (généré par npm run build) s'appelle "build"
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Erreur Getwd: %v", err)
	}
	staticDir := filepath.Join(workDir, "build")

	// Mappez directement les sous-dossiers critiques pour que les fichiers JS/CSS soient servis correctement
	fs := http.FileServer(http.Dir(staticDir))
	r.Handle("/_app/*", fs)
	r.Handle("/style/*", fs)
	r.Handle("/js/*", fs)
	r.Handle("/fontawesome/*", fs)
	r.Handle("/favicon.ico", fs)
	r.Handle("/favicon.png", fs)

	// Pour toutes les autres routes non définies, servir index.html (fallback SPA)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
	})

	// Lancer le serveur sur le port défini (par défaut 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("API & Front server running on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func runConsumer() {
	// Démarrer le consumer d'événements (par exemple, pour s'abonner sur "events.new")
	if err := consumer.StartConsumer(); err != nil {
		log.Fatalf("Erreur lors du démarrage du consumer : %v", err)
	}
}

func main() {
	// Initialiser la connexion NATS
	helpers.InitNats()

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
