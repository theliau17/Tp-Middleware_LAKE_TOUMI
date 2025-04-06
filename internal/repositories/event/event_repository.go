package event

import (
	"database/sql"
	"time"

	"middleware/example/internal/models"
)

// Repository définit l'interface pour les opérations sur les événements.
type Repository interface {
	// GetByUID récupère un événement en fonction de son UID.
	GetByUID(uid string) (*models.Event, error)
	// Upsert insère un nouvel événement ou met à jour un événement existant.
	Upsert(evt models.Event) error
}

// repository est la structure concrète implémentant Repository.
type repository struct {
	db *sql.DB
}

// NewRepository crée et retourne une nouvelle instance de repository.
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// GetByUID récupère un événement via sa clé UID.
// Si aucun événement n'est trouvé, nil est retourné sans erreur.
func (r *repository) GetByUID(uid string) (*models.Event, error) {
	query := `
		SELECT uid, summary, location, start_time, end_time
		FROM events
		WHERE uid = ?
	`
	row := r.db.QueryRow(query, uid)

	var evt models.Event
	err := row.Scan(&evt.UID, &evt.Summary, &evt.Location, &evt.StartTime, &evt.EndTime)
	if err != nil {
		if err == sql.ErrNoRows {
			// Aucun événement trouvé pour cet UID.
			return nil, nil
		}
		return nil, err
	}
	return &evt, nil
}

// Upsert insère ou met à jour un événement en utilisant "INSERT OR REPLACE".
// La colonne updated_at est mise à jour avec l'heure actuelle (UTC).
func (r *repository) Upsert(evt models.Event) error {
	query := `
		INSERT OR REPLACE INTO events (uid, summary, location, start_time, end_time, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query,
		evt.UID,
		evt.Summary,
		evt.Location,
		evt.StartTime,
		evt.EndTime,
		time.Now().UTC(),
	)
	return err
}
