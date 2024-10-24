package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type MessageService interface {
	GetMessages(
		ctx context.Context,
		channelId uuid.UUID,
		page, limit int,
	) ([]types.Message, error)

	GetMessageById(ctx context.Context, messageId uuid.UUID) (*types.Message, error)
	GetMessageReplies(ctx context.Context, messageId uuid.UUID) ([]types.Message, error)

	CreateMessage(ctx context.Context, dto types.CreateMessageDto) error
	UpdateMessage(
		ctx context.Context,
		messageId uuid.UUID,
		dto types.UpdateMessageDto,
	) error
	DeleteMessage(ctx context.Context, messageId uuid.UUID) error
}
