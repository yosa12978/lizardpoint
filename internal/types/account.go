package types

import (
	"context"
	"strings"
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

type CreateAccountDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//PasswordConfirm string `json:"password_confirm"`
}

func (c CreateAccountDto) Validate(ctx context.Context) (CreateAccountDto, map[string]string, bool) {
	problems := make(map[string]string)
	c.Username = strings.TrimSpace(c.Username)
	c.Password = strings.TrimSpace(c.Password)
	if len(c.Password) < 8 {
		problems["password"] = "password length must be >= 8"
	}
	if c.Username == "" {
		problems["username"] = "username can't be empty"
	}
	return c, problems, len(problems) == 0
}
