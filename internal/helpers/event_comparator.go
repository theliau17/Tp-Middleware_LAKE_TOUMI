package helpers

import (
	"middleware/example/internal/models"
)

func HasChanged(old, new models.Event) bool {
	// Comparaison du résumé
	if old.Summary != new.Summary {
		return true
	}

	// Comparaison de la localisation
	if old.Location != new.Location {
		return true
	}

	// Comparaison de la date/heure de début
	if !old.StartTime.Equal(new.StartTime) {
		return true
	}

	// Comparaison de la date/heure de fin
	if !old.EndTime.Equal(new.EndTime) {
		return true
	}

	// Si aucune des différences n'est détectée, retourner false
	return false
}
