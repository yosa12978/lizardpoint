package repos

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type accountPostgres struct {
	db     *sql.DB
	logger logging.Logger
}

func NewAccountPostgres(db *sql.DB, logger logging.Logger) AccountRepo {
	return &accountPostgres{
		db:     db,
		logger: logger,
	}
}

func (a *accountPostgres) GetAll(ctx context.Context) []types.Account {
	panic("unimplemented")
}

func (a *accountPostgres) GetById(ctx context.Context, id uuid.UUID) (*types.Account, error) {
	panic("unimplemented")
}

func (a *accountPostgres) GetByUsername(ctx context.Context, username string) (*types.Account, error) {
	panic("unimplemented")
}

func (a *accountPostgres) Create(ctx context.Context, account types.Account) error {
	panic("unimplemented")
}

func (a *accountPostgres) Update(ctx context.Context, account types.Account) error {
	panic("unimplemented")
}

func (a *accountPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
