package domain

import "time"

// ImportRecord represents a row imported from an Excel file
type ImportRecord struct {
	ID        string                 `json:"id"`
	Data      map[string]any `json:"data,omitempty"`
	ProfileID *string                `json:"profile_id,omitempty"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
