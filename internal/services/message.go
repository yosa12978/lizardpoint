package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/repos"
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

	CreateMessage(
		ctx context.Context,
		accountId, channelId uuid.UUID,
		message string,
	) error
	CreateReply(
		ctx context.Context,
		accountId,
		channelId,
		parentId uuid.UUID,
		content string,
	) error
	UpdateMessage(
		ctx context.Context,
		accountId,
		messageId uuid.UUID,
		content string,
	) error
	DeleteMessage(
		ctx context.Context,
		accountId,
		messageId uuid.UUID,
	) error
}

type messageService struct {
	messageRepo repos.MessageRepo
	logger      logging.Logger
}

func NewMessageService(
	messageRepo repos.MessageRepo,
	logger logging.Logger,
) MessageService {
	return &messageService{
		messageRepo: messageRepo,
		logger:      logger,
	}
}

// CreateMessage implements MessageService.
func (m *messageService) CreateMessage(ctx context.Context, accountId uuid.UUID, channelId uuid.UUID, message string) error {
	panic("unimplemented")
}

// CreateReply implements MessageService.
func (m *messageService) CreateReply(ctx context.Context, accountId uuid.UUID, channelId uuid.UUID, parentId uuid.UUID, content string) error {
	panic("unimplemented")
}

// DeleteMessage implements MessageService.
func (m *messageService) DeleteMessage(ctx context.Context, accountId uuid.UUID, messageId uuid.UUID) error {
	panic("unimplemented")
}

// GetMessageById implements MessageService.
func (m *messageService) GetMessageById(ctx context.Context, messageId uuid.UUID) (*types.Message, error) {
	panic("unimplemented")
}

// GetMessageReplies implements MessageService.
func (m *messageService) GetMessageReplies(ctx context.Context, messageId uuid.UUID) ([]types.Message, error) {
	panic("unimplemented")
}

// GetMessages implements MessageService.
func (m *messageService) GetMessages(ctx context.Context, channelId uuid.UUID, page int, limit int) ([]types.Message, error) {
	panic("unimplemented")
}

// UpdateMessage implements MessageService.
func (m *messageService) UpdateMessage(ctx context.Context, accountId uuid.UUID, messageId uuid.UUID, content string) error {
	panic("unimplemented")
}
