package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"          db:"id"`
	Login     string    `json:"login"       db:"login"`
	Email     string    `json:"email"       db:"email"`
	CreatedAt time.Time `json:"created_at"  db:"created_at"`
}
