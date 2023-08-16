package data

import (
	"time"

	"github.com/google/uuid"
)

type Checklist struct {
	id        uuid.UUID // PK
	taskId    uuid.UUID // FK
	name      string
	isChecked bool
	position  int
	createdAt time.Time
	createdBy uuid.UUID // user id
}
