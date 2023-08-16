package data

import (
	"github.com/google/uuid"
)

type CardLabel struct {
	labelId uuid.UUID // FK
	taskId  uuid.UUID // FK
}
