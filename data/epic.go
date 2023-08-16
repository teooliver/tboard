package data

import (
	"time"

	"github.com/google/uuid"
)

type Epic struct {
	id          uuid.UUID
	title       string
	description string
	createdAt   time.Time
	createdBy   uuid.UUID
}

func NewEpic(title string, description string, createdBy uuid.UUID) *Epic {
	return &Epic{
		id:          uuid.New(),
		title:       title,
		description: description,
		createdAt:   time.Now(),
		createdBy:   createdBy,
	}
}
