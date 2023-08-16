package data

import (
	"time"

	"github.com/google/uuid"
)

type ChecklistItem struct {
	id        uuid.UUID // PK
	taskId    uuid.UUID // FK
	name      string    // name of the item
	isChecked bool
	position  int
	createdAt time.Time
	createdBy uuid.UUID // user id
}
