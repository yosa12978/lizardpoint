package types

import (
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
