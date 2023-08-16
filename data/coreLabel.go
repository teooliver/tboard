package data

import (
	"time"

	"github.com/google/uuid"
)

type CoreLabel struct {
	id        uuid.UUID // PK
	name      string
	color     string
	createdAt time.Time
	createdBy uuid.UUID // user id
}
