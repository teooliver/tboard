package data

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id        uuid.UUID // uuid
	name      string
	email     string
	createdAt time.Time
}

func NewUser(name string, email string) *User {
	return &User{
		id:        uuid.New(),
		name:      name,
		email:     email,
		createdAt: time.Now(),
	}
}
