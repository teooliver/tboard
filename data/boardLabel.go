package data

import (
	"time"

	"github.com/google/uuid"
)

type BoardLabel struct {
	id        uuid.UUID // PK
	boardId   uuid.UUID // FK
	name      string
	color     string
	createdAt time.Time
	createdBy uuid.UUID // user id
}
