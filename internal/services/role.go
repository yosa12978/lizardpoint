package services

import (
	"context"

	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/repos"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type RoleService interface {
	GetRoles(ctx context.Context) ([]types.Role, error)
	GetRoleByName(ctx context.Context, name string) (*types.Role, error)

	CreateRole(ctx context.Context, name string) error
	UpdateRole(ctx context.Context, oldName string, newName string) error
	DeleteRole(ctx context.Context, role string) error
}

type roleService struct {
	roleRepo repos.RoleRepo
	logger   logging.Logger
}

func NewRoleService(
	roleRepo repos.RoleRepo,
	logger logging.Logger,
) RoleService {
	return &roleService{
		roleRepo: roleRepo,
		logger:   logger,
	}
}

// CreateRole implements RoleService.
func (r *roleService) CreateRole(ctx context.Context, name string) error {
	panic("unimplemented")
}

// DeleteRole implements RoleService.
func (r *roleService) DeleteRole(ctx context.Context, role string) error {
	panic("unimplemented")
}

// GetRoleByName implements RoleService.
func (r *roleService) GetRoleByName(ctx context.Context, name string) (*types.Role, error) {
	panic("unimplemented")
}

// GetRoles implements RoleService.
func (r *roleService) GetRoles(ctx context.Context) ([]types.Role, error) {
	panic("unimplemented")
}

// UpdateRole implements RoleService.
func (r *roleService) UpdateRole(ctx context.Context, oldName string, newName string) error {
	panic("unimplemented")
}
