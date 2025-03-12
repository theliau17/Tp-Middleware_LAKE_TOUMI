package timetable

import (
	"os"

	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

// Service gère la logique de récupération des événements ICS
type Service interface {
	GetAllEvents() ([]models.Event, error)
}

type service struct {
	icsPath string
}

// NewService crée un nouveau service Timetable
func NewService(icsPath string) Service {
	return &service{icsPath: icsPath}
}

// GetAllEvents lit le fichier ICS, le parse et renvoie les événements
func (s *service) GetAllEvents() ([]models.Event, error) {
	file, err := os.Open(s.icsPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	events, err := helpers.ParseICS(file)
	if err != nil {
		return nil, err
	}
	return events, nil
}
