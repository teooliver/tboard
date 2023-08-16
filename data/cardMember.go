package data

import (
	"time"

	"github.com/google/uuid"
)

type CardMember struct {
	userId  uuid.UUID // FK
	taskId  uuid.UUID // FK
	addedAt time.Time
	addedBy uuid.UUID // user id
}
