package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/repos"
	"github.com/yosa12978/lizardpoint/internal/types"
	"github.com/yosa12978/lizardpoint/pkg/utils"
)

type AccountService interface {
	GetAccounts(ctx context.Context) ([]types.Account, error)
	GetAccountById(ctx context.Context, accountId uuid.UUID) (*types.Account, error)

	SetActive(ctx context.Context, accountId uuid.UUID) error
	SetInactive(ctx context.Context, accountId uuid.UUID) error

	CreateAccount(ctx context.Context, username, password string, defaultRole *types.Role) error
	ChangePassword(ctx context.Context, accountId uuid.UUID, oldPassword, newPassword string) error
	DeleteAccount(ctx context.Context, accountId uuid.UUID) error
}

type accountService struct {
	accountRepo repos.AccountRepo
	roleRepo    repos.RoleRepo
	logger      logging.Logger
}

func NewAccountService(
	accountRepo repos.AccountRepo,
	roleRepo repos.RoleRepo,
	logger logging.Logger,
) AccountService {
	return &accountService{
		accountRepo: accountRepo,
		roleRepo:    roleRepo,
		logger:      logger,
	}
}

func (a *accountService) SetActive(ctx context.Context, accountId uuid.UUID) error {
	return nil
}

func (a *accountService) SetInactive(ctx context.Context, accountId uuid.UUID) error {
	return nil
}

// ChangePassword implements AccountService.
func (a *accountService) ChangePassword(ctx context.Context, accountId uuid.UUID, oldPassword string, newPassword string) error {
	account, err := a.accountRepo.GetById(ctx, accountId)
	if err != nil {
		return err
	}
	salt := uuid.NewString()
	passwordHash, err := utils.HashPassword(newPassword + salt)
	if err != nil {
		return types.NewErrInternalFailure(err)
	}
	account.PasswordHash = passwordHash
	account.Salt = salt
	return a.accountRepo.Update(ctx, *account)
}

// CreateAccount implements AccountService.
func (a *accountService) CreateAccount(ctx context.Context, username, password string, defaultRole *types.Role) error {
	salt := uuid.NewString()
	passwordHash, err := utils.HashPassword(password)
	if err != nil {
		return types.NewErrInternalFailure(err)
	}
	account := types.Account{
		Id:           uuid.New(),
		Username:     username,
		PasswordHash: passwordHash,
		Salt:         salt,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	// this looks shitty
	if defaultRole != nil {
		return a.accountRepo.CreateWithDefaultRole(ctx, account, *defaultRole)
	}
	return a.accountRepo.Create(ctx, account)
}

// DeleteAccount implements AccountService.
func (a *accountService) DeleteAccount(ctx context.Context, accountId uuid.UUID) error {
	return a.accountRepo.Delete(ctx, accountId)
}

// GetAccountById implements AccountService.
func (a *accountService) GetAccountById(ctx context.Context, accountId uuid.UUID) (*types.Account, error) {
	return a.accountRepo.GetById(ctx, accountId)
}

// GetAccounts implements AccountService.
func (a *accountService) GetAccounts(ctx context.Context) ([]types.Account, error) {
	return a.accountRepo.GetAll(ctx)
}
