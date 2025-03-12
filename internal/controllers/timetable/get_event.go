package timetable

import (
	"encoding/json"
	"net/http"

	"middleware/example/internal/models"
	"middleware/example/internal/services/timetable"
)

// GetEvents gère la route GET /events
// GetEvents godoc
// @Summary Récupère tous les événements
// @Description Lit un fichier ICS et renvoie la liste des événements
// @Tags Timetable
// @Produce json
// @Success 200 {array} models.Event
// @Failure 500 {object} models.APIError
// @Router /events [get]
func GetEvents(svc timetable.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := svc.GetAllEvents()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// Utilise le modèle APIError défini dans internal/models/errors.go
			json.NewEncoder(w).Encode(models.APIError{Message: "Unable to get events"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	}
}
