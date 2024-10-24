package types

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

type Channel struct {
	Id               uuid.UUID
	Name             string
	ReadPermissions  []Role
	WritePermissions []Role
}

type CreateChannelDto struct {
	Name string `json:"name"`
}

func (c CreateChannelDto) Validate(ctx context.Context) (
	CreateChannelDto, map[string]string, bool,
) {
	problems := make(map[string]string)
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		problems["name"] = "name is required"
	}
	return c, problems, len(problems) == 0
}

type UpdateChannelDto struct {
	Name string `json:"name"`
}

func (c UpdateChannelDto) Validate(ctx context.Context) (
	UpdateChannelDto, map[string]string, bool,
) {
	problems := make(map[string]string)
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		problems["name"] = "name can't be empty"
	}
	return c, problems, len(problems) == 0
}
