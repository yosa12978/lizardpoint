package repos

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
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

var getAllAccountsSQL = `
	SELECT a.username, a.password_hash, a.is_active, a.created_at, a.updated_at,
	ARRAY(SELECT role_name FROM accounts_roles ar WHERE ar.account_id = a.id) AS roles 
	FROM accounts a;
`

func (a *accountPostgres) GetAll(ctx context.Context) ([]types.Account, error) {
	a.logger.Debug("fetching all accounts")
	var accounts []types.Account
	row, err := a.db.QueryContext(ctx, getAllAccountsSQL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return accounts, nil
		}
		return accounts, types.NewErrInternalFailure(err)
	}
	for row.Next() {
		account := types.Account{}
		var rolesStr []string
		row.Scan(
			&account.Username,
			&account.PasswordHash,
			&account.IsActive,
			&account.CreatedAt,
			&account.UpdatedAt,
			(*pq.StringArray)(&rolesStr),
		)
		for _, name := range rolesStr {
			account.Roles = append(account.Roles, types.Role{Name: name})
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

var getAccountByIdSQL = `
	SELECT a.username, a.password_hash, a.is_active, a.created_at, a.updated_at,
	ARRAY(SELECT role_name FROM accounts_roles ar WHERE ar.account_id = a.id) AS roles 
	FROM accounts a WHERE a.id = $1;
`

func (a *accountPostgres) GetById(ctx context.Context, id uuid.UUID) (*types.Account, error) {
	a.logger.Debug("fetching account by id", "id", id.String())
	var account types.Account
	var rolesStr []string
	err := a.db.QueryRowContext(ctx, getAccountByIdSQL, id).
		Scan(
			&account.Username,
			&account.PasswordHash,
			&account.IsActive,
			&account.CreatedAt,
			&account.UpdatedAt,
			(*pq.StringArray)(&rolesStr),
		)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.NewErrNotFound(err)
		}
		return nil, types.NewErrInternalFailure(err)
	}

	for _, name := range rolesStr {
		account.Roles = append(account.Roles, types.Role{Name: name})
	}
	return &account, err
}

var getAccountByUsernameSQL = `
	SELECT a.username, a.password_hash, a.is_active, a.created_at, a.updated_at,
	ARRAY(SELECT role_name FROM accounts_roles ar WHERE ar.account_id = a.id) AS roles 
	FROM accounts a WHERE a.username = $1;
`

func (a *accountPostgres) GetByUsername(ctx context.Context, username string) (*types.Account, error) {
	a.logger.Debug("fetching account by username", "username", username)
	var account types.Account
	var rolesStr []string
	err := a.db.QueryRowContext(ctx, getAccountByUsernameSQL, username).
		Scan(
			&account.Username,
			&account.PasswordHash,
			&account.IsActive,
			&account.CreatedAt,
			&account.UpdatedAt,
			(*pq.StringArray)(&rolesStr),
		)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.NewErrNotFound(err)
		}
		return nil, types.NewErrInternalFailure(err)
	}

	for _, name := range rolesStr {
		account.Roles = append(account.Roles, types.Role{Name: name})
	}
	return &account, nil
}

var insertAccountSQL = `
	INSERT INTO accounts 
		(username, password_hash, is_active, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5);
`

func (a *accountPostgres) Create(ctx context.Context, account types.Account) error {
	a.logger.Debug("creating new account", "username", account.Username)
	_, err := a.db.ExecContext(ctx,
		insertAccountSQL,
		account.Username,
		account.PasswordHash,
		account.IsActive,
		account.CreatedAt,
		account.UpdatedAt,
	)
	if err != nil {
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var updateAccountSQL = `
	UPDATE accounts SET username=$1, password_hash=$2, is_active=$3, updated_at=$4
	WHERE id=$5
`

func (a *accountPostgres) Update(ctx context.Context, account types.Account) error {
	a.logger.Debug("updating account", "username", account.Id.String())
	_, err := a.db.ExecContext(ctx,
		updateAccountSQL,
		account.Username,
		account.PasswordHash,
		account.IsActive,
		time.Now().UTC(), //account.UpdatedAt,
		account.Id,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.NewErrNotFound(err)
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var deleteAccountSQL = `
	DELETE FROM accounts WHERE id=$1;
`

func (a *accountPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	a.logger.Debug("removing account", "id", id.String())
	_, err := a.db.ExecContext(ctx, deleteAccountSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.NewErrNotFound(err)
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}
