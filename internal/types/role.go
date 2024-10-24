package types

import (
	"context"
	"strings"
)

type Role struct {
	Name string
}

type CreateRoleDto struct {
	Name string `json:"name"`
}

func (c CreateRoleDto) Validate(ctx context.Context) (
	CreateRoleDto, map[string]string, bool,
) {
	problems := make(map[string]string)
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		problems["name"] = "name can't be empty"
	}
	return c, problems, len(problems) == 0
}

type UpdateRoleDto struct {
	Name string `json:"name"`
}

func (c UpdateRoleDto) Validate(ctx context.Context) (
	UpdateRoleDto, map[string]string, bool,
) {
	problems := make(map[string]string)
	c.Name = strings.TrimSpace(c.Name)
	if c.Name == "" {
		problems["name"] = "name can't be empty"
	}
	return c, problems, len(problems) == 0
}
