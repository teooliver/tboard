package data

import (
	"time"

	"github.com/google/uuid"
)

type Column struct {
	id        uuid.UUID
	name      string
	position  int
	tasks     []uuid.UUID
	createdAt time.Time
	createdBy uuid.UUID
}

func NewColumn(name string, createdBy uuid.UUID) *Column {
	return &Column{
		id:        uuid.New(),
		name:      name,
		createdAt: time.Now(),
		createdBy: createdBy,
	}
}
