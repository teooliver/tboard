package data

import (
	"time"

	"github.com/google/uuid"
)

type TaskActivity struct {
	id        uuid.UUID // PK
	taskId    uuid.UUID // FK
	userId    uuid.UUID // FK
	activity  string
	createdAt time.Time
	createdBy uuid.UUID // user id
}
