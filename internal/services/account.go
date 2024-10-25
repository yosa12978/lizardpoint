package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/repos"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type AccountService interface {
	GetAccounts(ctx context.Context, page, limit int) ([]types.Account, error)
	GetAccountById(ctx context.Context) (*types.Account, error)

	CreateAccount(ctx context.Context, username, password string) error
	ChangePassword(ctx context.Context, accountId uuid.UUID, oldPassword, newPassword string) error
	DeleteAccount(ctx context.Context, accountId uuid.UUID) error
}

type accountService struct {
	accountRepo repos.AccountRepo
	logger      logging.Logger
}

func NewAccountService(
	accountRepo repos.AccountRepo,
	logger logging.Logger,
) AccountService {
	return &accountService{
		accountRepo: accountRepo,
		logger:      logger,
	}
}

// ChangePassword implements AccountService.
func (a *accountService) ChangePassword(ctx context.Context, accountId uuid.UUID, oldPassword string, newPassword string) error {
	panic("unimplemented")
}

// CreateAccount implements AccountService.
func (a *accountService) CreateAccount(ctx context.Context, username string, password string) error {
	panic("unimplemented")
}

// DeleteAccount implements AccountService.
func (a *accountService) DeleteAccount(ctx context.Context, accountId uuid.UUID) error {
	panic("unimplemented")
}

// GetAccountById implements AccountService.
func (a *accountService) GetAccountById(ctx context.Context) (*types.Account, error) {
	panic("unimplemented")
}

// GetAccounts implements AccountService.
func (a *accountService) GetAccounts(ctx context.Context, page int, limit int) ([]types.Account, error) {
	panic("unimplemented")
}
