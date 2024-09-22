package types

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id           uuid.UUID
	Username     string
	PasswordHash string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Roles        []Role
}
