package repos

import (
	"context"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type ChannelRepo interface {
	GetAll(ctx context.Context) ([]types.Channel, error)
	GetById(ctx context.Context, id uuid.UUID) (*types.Channel, error)

	Create(ctx context.Context, channel types.Channel) error
	Update(ctx context.Context, channel types.Channel) error
	Delete(ctx context.Context, id uuid.UUID) error
}
