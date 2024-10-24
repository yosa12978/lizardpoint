package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type ChannelService interface {
	GetChannels(ctx context.Context) ([]types.Channel, error)
	GetChannelById(ctx context.Context, channelId uuid.UUID) (*types.Channel, error)

	AddWritePermission(ctx context.Context, channelId uuid.UUID, role string) error
	AddReadPermission(ctx context.Context, channelId uuid.UUID, role string) error
	RemoveWritePermission(ctx context.Context, channelId uuid.UUID, role string) error
	RemoveReadPermission(ctx context.Context, channelId uuid.UUID, role string) error

	CreateChannel(ctx context.Context, dto types.CreateChannelDto) error
	UpdateChannel(ctx context.Context, channelId uuid.UUID, dto types.UpdateChannelDto) error
	DeleteChannel(ctx context.Context, channelId uuid.UUID) error
}
