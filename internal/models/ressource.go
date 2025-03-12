package models

import "time"

type Resource struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	FilePath  string    `json:"file_path"`
	CreatedAt time.Time `json:"created_at"`
}
