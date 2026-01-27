package store

import (
	"time"

	"github.com/google/uuid"
)

// Both times store the timestamps as minutes since midnight
type Entry struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   string    `json:"project_id"`
	Date        string    `json:"date"`
	Description string    `json:"description"`
	StartTime   Minutes   `json:"start_time"`
	EndTime     Minutes   `json:"end_time"`
	CreatedAt   time.Time `json:"created_at"`
}

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
