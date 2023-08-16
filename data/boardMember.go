package data

import (
	"time"

	"github.com/google/uuid"
)

// Join Table
type BoardMember struct {
	user_id  uuid.UUID // FK
	board_id uuid.UUID // FK

	createdAt time.Time
	createdBy uuid.UUID // user id
}
