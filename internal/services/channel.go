package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/repos"
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

type channelService struct {
	channelRepo repos.ChannelRepo
	logger      logging.Logger
}

func NewChannelService(
	channelRepo repos.ChannelRepo,
	logger logging.Logger,
) ChannelService {
	return &channelService{
		channelRepo: channelRepo,
		logger:      logger,
	}
}

// AddReadPermission implements ChannelService.
func (c *channelService) AddReadPermission(ctx context.Context, channelId uuid.UUID, role string) error {
	panic("unimplemented")
}

// AddWritePermission implements ChannelService.
func (c *channelService) AddWritePermission(ctx context.Context, channelId uuid.UUID, role string) error {
	panic("unimplemented")
}

// CreateChannel implements ChannelService.
func (c *channelService) CreateChannel(ctx context.Context, dto types.CreateChannelDto) error {
	panic("unimplemented")
}

// DeleteChannel implements ChannelService.
func (c *channelService) DeleteChannel(ctx context.Context, channelId uuid.UUID) error {
	panic("unimplemented")
}

// GetChannelById implements ChannelService.
func (c *channelService) GetChannelById(ctx context.Context, channelId uuid.UUID) (*types.Channel, error) {
	panic("unimplemented")
}

// GetChannels implements ChannelService.
func (c *channelService) GetChannels(ctx context.Context) ([]types.Channel, error) {
	panic("unimplemented")
}

// RemoveReadPermission implements ChannelService.
func (c *channelService) RemoveReadPermission(ctx context.Context, channelId uuid.UUID, role string) error {
	panic("unimplemented")
}

// RemoveWritePermission implements ChannelService.
func (c *channelService) RemoveWritePermission(ctx context.Context, channelId uuid.UUID, role string) error {
	panic("unimplemented")
}

// UpdateChannel implements ChannelService.
func (c *channelService) UpdateChannel(ctx context.Context, channelId uuid.UUID, dto types.UpdateChannelDto) error {
	panic("unimplemented")
}
