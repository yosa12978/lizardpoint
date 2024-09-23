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

var getAllAccountsSQL = `
	SELECT a.username, a.password_hash, a.is_active, a.created_at, a.updated_at,
	ARRAY(SELECT role_name FROM accounts_roles ar WHERE ar.account_id = a.id) AS roles 
	FROM accounts a;
`

func (a *accountPostgres) GetAll(ctx context.Context) ([]types.Account, error) {
	var accounts []types.Account
	row, err := a.db.QueryContext(ctx, getAllAccountsSQL)
	if err != nil {
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
			&rolesStr,
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
	panic("unimplemented")
}

var getAccountByUsernameSQL = `
	SELECT a.username, a.password_hash, a.is_active, a.created_at, a.updated_at,
	ARRAY(SELECT role_name FROM accounts_roles ar WHERE ar.account_id = a.id) AS roles 
	FROM accounts a WHERE a.username = $1;
`

func (a *accountPostgres) GetByUsername(ctx context.Context, username string) (*types.Account, error) {
	panic("unimplemented")
}

var insertAccountSQL = `
	INSERT INTO accounts 
		(username, password_hash, is_active, created_at, updated_at) 
	VALUES ($1, $2, $3, $4, $5);
`

func (a *accountPostgres) Create(ctx context.Context, account types.Account) error {
	panic("unimplemented")
}

var updateAccountSQL = `
	UPDATE accounts SET username=$1, password_hash=$2, is_active=$3, updated_at=$4
	WHERE id=$5
`

func (a *accountPostgres) Update(ctx context.Context, account types.Account) error {
	panic("unimplemented")
}

var deleteAccountSQL = `
	DELETE FROM accounts WHERE id=$1;
`

func (a *accountPostgres) Delete(ctx context.Context, id uuid.UUID) error {
	panic("unimplemented")
}
