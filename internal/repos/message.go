package repos

import (
	"context"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type Message interface {
	GetByChannelId(ctx context.Context, channelId uuid.UUID, page int, limit int) ([]types.Message, error)
	GetReplies(ctx context.Context, parentId uuid.UUID, page int, limit int) ([]types.Message, error)

	Create(ctx context.Context, message types.Message) error
	Update(ctx context.Context, message types.Message) error
	Delete(ctx context.Context, id uuid.UUID) error
}
