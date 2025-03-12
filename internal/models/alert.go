package models

import "time"

type Alert struct {
	ID         string   `json:"id"`
	Recipients []string `json:"recipients"` // destinataires
	ResourceID string   `json:"resource_id"`
	// ex: "always", "room_changed_only", etc.
	Condition string    `json:"condition"`
	CreatedAt time.Time `json:"created_at"`
}
