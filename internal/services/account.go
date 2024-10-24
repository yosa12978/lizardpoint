package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type AccountService interface {
	GetAccounts(ctx context.Context, page, limit int) ([]types.Account, error)
	GetAccountById(ctx context.Context) (*types.Account, error)

	CreateAccount(ctx context.Context, dto types.CreateAccountDto) error
	ChangePassword(ctx context.Context, accountId uuid.UUID, dto types.UpdatePasswordDto) error
	DeleteAccount(ctx context.Context, accountId uuid.UUID) error
}
