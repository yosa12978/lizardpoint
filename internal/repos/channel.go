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

	AddWritePermission(ctx context.Context, channelId uuid.UUID, role string) error
	AddReadPermission(ctx context.Context, channelId uuid.UUID, role string) error
	RemoveWritePermission(ctx context.Context, channelId uuid.UUID, role string) error
	RemoveReadPermission(ctx context.Context, channelId uuid.UUID, role string) error
}
