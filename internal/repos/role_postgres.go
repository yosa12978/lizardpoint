package repos

import (
	"context"
	"database/sql"
	"errors"

	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/types"
)

type rolePostgres struct {
	db     *sql.DB
	logger logging.Logger
}

func NewRolePostgres(db *sql.DB, logger logging.Logger) RoleRepo {
	return &rolePostgres{
		db:     db,
		logger: logger,
	}
}

var getAllRolesSQL = `
	SELECT name FROM roles;
`

func (r *rolePostgres) GetAll(ctx context.Context) ([]types.Role, error) {
	roles := []types.Role{}
	row, err := r.db.QueryContext(ctx, getAllRolesSQL)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return roles, nil
		}
		return roles, types.NewErrInternalFailure(err)
	}
	for row.Next() {
		role := types.Role{}
		row.Scan(&role.Name)
		roles = append(roles, role)
	}
	return roles, nil
}

var getRoleByIdSQL = `
	SELECT name FROM roles WHERE name=$1;
`

func (r *rolePostgres) GetByName(ctx context.Context, name string) (*types.Role, error) {
	role := types.Role{}
	err := r.db.QueryRowContext(ctx, getRoleByIdSQL, name).Scan(&role.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.ErrNotFound
		}
		return nil, types.NewErrInternalFailure(err)
	}
	return &role, nil
}

var insertRoleSQL = `
	INSERT INTO roles (name) VALUES ($1);
`

func (r *rolePostgres) Create(ctx context.Context, role types.Role) error {
	_, err := r.db.ExecContext(ctx, insertRoleSQL, role.Name)
	if err != nil {
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var updateRoleSQL = `
	UPDATE roles SET name=$1 WHERE name=$2;
`

func (r *rolePostgres) Update(ctx context.Context, oldName, newName string) error {
	_, err := r.db.ExecContext(ctx, updateRoleSQL, newName, oldName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.ErrNotFound
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}

var deleteRoleSQL = `
	DELETE FROM roles WHERE name=$1;
`

func (r *rolePostgres) Delete(ctx context.Context, name string) error {
	_, err := r.db.ExecContext(ctx, deleteRoleSQL, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return types.ErrNotFound
		}
		return types.NewErrInternalFailure(err)
	}
	return nil
}
