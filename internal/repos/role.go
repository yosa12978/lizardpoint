package repos

import (
	"context"

	"github.com/yosa12978/lizardpoint/internal/types"
)

type RoleRepo interface {
	GetAll(ctx context.Context) ([]types.Role, error)
	GetByName(ctx context.Context, name string) (*types.Role, error)

	Create(ctx context.Context, role types.Role) error
	Update(ctx context.Context, oldName, newName string) error
	Delete(ctx context.Context, name string) error
}
