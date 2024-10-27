package repos

import (
	"context"
	"database/sql"
)

func runInTx(ctx context.Context, db *sql.DB, f func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := f(tx); err == nil {
		return tx.Commit()
	}
	if err := tx.Rollback(); err != nil {
		return err
	}
	return err
}
