package models

import "time"

// Event représente un événement issu d'un fichier ICS.
type Event struct {
	UID       string    `json:"uid"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Summary   string    `json:"summary"`
	Location  string    `json:"location"`
}
