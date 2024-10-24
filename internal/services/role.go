package services

import (
	"context"

	"github.com/yosa12978/lizardpoint/internal/types"
)

type RoleService interface {
	GetRoles(ctx context.Context) ([]types.Role, error)
	GetRoleByName(ctx context.Context, name string) (*types.Role, error)

	CreateRole(ctx context.Context, dto types.CreateRoleDto) error
	UpdateRole(ctx context.Context, role string, dto types.UpdateRoleDto) error
	DeleteRole(ctx context.Context, role string) error
}
