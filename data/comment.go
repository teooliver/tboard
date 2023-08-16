package data

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	id        uuid.UUID // PK
	taskId    uuid.UUID // FK
	comment   string
	createdAt time.Time
	createdBy uuid.UUID // FK user id
}
