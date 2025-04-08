package domain

import (
	"time"
)

// User represents a system user.
// Prefer using NewUser(...) to ensure required fields are set.
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
