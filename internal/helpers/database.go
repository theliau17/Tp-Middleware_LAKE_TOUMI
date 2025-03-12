package helpers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// OpenDB ouvre la base SQLite avec le DSN donné (ex: "file:collections.db").
func OpenDB(dsn string) (*sql.DB, error) {
	return sql.Open("sqlite3", dsn)
}

// InitSchema crée la table resources si elle n'existe pas.
// Tu pourras plus tard ajouter la table alerts.
func InitSchema(db *sql.DB) error {
	// Exemple minimal pour la table "resources".
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS resources (
        id TEXT PRIMARY KEY,
        name TEXT,
        file_path TEXT,
        created_at DATETIME
    );
    `)
	return err
}
