package types

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Id                    uuid.UUID
	Content               string
	Edited                bool
	CreatedAt             time.Time
	UpdatedAt             time.Time
	ChannelId             uuid.UUID
	ChannelName           string
	AccountId             uuid.UUID
	AccountUsername       string
	ParentId              uuid.UUID
	ParentAccountId       uuid.UUID
	ParentAccountUsername string
}

type CreateMessageDto struct {
	Content  string    `json:"content"`
	ParentId uuid.UUID `json:"parent_id,omitempty"`
}

func (c CreateMessageDto) Validate(ctx context.Context) (
	CreateMessageDto, map[string]string, bool,
) {
	problems := make(map[string]string)
	c.Content = strings.TrimSpace(c.Content)
	if c.Content == "" {
		problems["content"] = "content can't be empty"
	}
	return c, problems, len(problems) == 0
}

type UpdateMessageDto struct {
	Content string `json:"content"`
}

func (c UpdateMessageDto) Validate(ctx context.Context) (
	UpdateMessageDto, map[string]string, bool,
) {
	problems := make(map[string]string)
	c.Content = strings.TrimSpace(c.Content)
	if c.Content == "" {
		problems["content"] = "content can't be empty"
	}
	return c, problems, len(problems) == 0
}
