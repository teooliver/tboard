package data

import (
	"time"

	"github.com/google/uuid"
)

type TaskAttachment struct {
	id           uuid.UUID // PK
	taskId       uuid.UUID // FK
	uploadedData time.Time // FK
	filename     string
	location     string
	createdAt    time.Time
	createdBy    uuid.UUID // user id
}
