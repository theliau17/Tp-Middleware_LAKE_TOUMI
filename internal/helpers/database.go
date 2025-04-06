package helpers

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDB(dsn string) (*sql.DB, error) {
	return sql.Open("sqlite3", dsn)
}

func InitSchema(db *sql.DB) error {
	// Création de la table resources
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS resources (
            id TEXT PRIMARY KEY,
            name TEXT,
            file_path TEXT,
            created_at DATETIME
        );
    `)
	if err != nil {
		return err
	}

	// Création de la table events
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS events (
            uid TEXT PRIMARY KEY,
            summary TEXT,
            location TEXT,
            start_time DATETIME,
            end_time DATETIME,
            updated_at DATETIME
        );
    `)
	if err != nil {
		return err
	}

	// Vous pouvez ajouter d'autres tables ici, par exemple alerts, etc.
	// _, err = db.Exec(`
	//     CREATE TABLE IF NOT EXISTS alerts (
	//         ...
	//     );
	// `)
	// if err != nil {
	//     return err
	// }

	return nil
}
