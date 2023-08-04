package data

import (
	"time"

	"github.com/google/uuid"
)

type Board struct {
	id         uuid.UUID // uuid
	name       string
	columnRefs []uuid.UUID // column ids
	createdAt  time.Time
	createdBy  uuid.UUID // user id
}

func NewBoard(name string, createdBy uuid.UUID) *Board {
	return &Board{
		id:         uuid.New(),
		name:       name,
		columnRefs: make([]uuid.UUID, 0),
		createdAt:  time.Now(),
		createdBy:  createdBy,
	}
}

func (b *Board) addColumn(columnRef uuid.UUID) {
	b.columnRefs = append(b.columnRefs, columnRef)
}
