package data

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	id           uuid.UUID
	title        string
	description  string
	estimation   int // hours
	dueData      time.Time
	remindedData time.Time
	assignedTo   uuid.UUID // user id
	epicRef      uuid.UUID // reference to the epic
	columnRef    uuid.UUID // column id
	isActive     bool      // archive
	createdAt    time.Time
	createdBy    uuid.UUID
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

func (t *Task) moveToEpic(epicRef uuid.UUID) {
	t.epicRef = epicRef
}

func (t *Task) assignTo(userID uuid.UUID) {
	t.assignedTo = userID
}
