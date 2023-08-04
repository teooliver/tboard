package data

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	id          uuid.UUID
	title       string
	description string
	estimation  int       // hours
	assignedTo  string    // user id
	sprintRef   string    // sprint id
	boardRef    string    // board id
	columnRef   uuid.UUID // column id
	createdAt   time.Time
	createdBy   uuid.UUID
}

func NewTask(title string, description string, createdBy uuid.UUID) *Task {
	return &Task{
		id:          uuid.New(),
		title:       title,
		description: description,
		createdAt:   time.Now(),
		createdBy:   createdBy,
	}
}

func (t *Task) moveToColumn(columnRef uuid.UUID) {
	t.columnRef = columnRef
}
