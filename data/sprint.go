package data

import (
	"time"

	"github.com/google/uuid"
)

var DefaultSprintSize int = 2 // two weeks

type Sprint struct {
	id        uuid.UUID
	name      string
	boardRefs []uuid.UUID // board ids
	size      int         // amount of weeks
	createdAt time.Time
	createdBy uuid.UUID // use id
}

func NewSprint(name string, createdBy uuid.UUID, size int) *Sprint {
	return &Sprint{
		id:        uuid.New(),
		name:      name,
		boardRefs: make([]uuid.UUID, 0),
		size:      size,
		createdAt: time.Now(),
		createdBy: createdBy,
	}
}

func (s *Sprint) addBoard(boardRef uuid.UUID) {
	s.boardRefs = append(s.boardRefs, boardRef)
}

// Function to "cache" Sprint in Normalized form so its quicker and easyer to display it
func (s *Sprint) cacheNormalizedData() {
	// todo()!
}

func (s *Sprint) ToJSON() {
	//  todo()!
}
