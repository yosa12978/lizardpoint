package repos

import (
	"context"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type AccountRepo interface {
	GetAll(ctx context.Context) ([]types.Account, error)
	GetById(ctx context.Context, id uuid.UUID) (*types.Account, error)
	GetByUsername(ctx context.Context, username string) (*types.Account, error)

	Create(ctx context.Context, account types.Account) error
	Update(ctx context.Context, account types.Account) error
	Delete(ctx context.Context, id uuid.UUID) error
}
