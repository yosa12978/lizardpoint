package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/yosa12978/lizardpoint/internal/config"
	"github.com/yosa12978/lizardpoint/internal/logging"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	pgInit sync.Once
	pgDb   *sql.DB
)

// return error instead of calling panic
func connectPostgres(ctx context.Context, logger logging.Logger) func() {
	return func() {
		conf := config.Get()
		url := fmt.Sprintf(
			"postgres://%s:%s@%s/%s?sslmode=%s",
			conf.Postgres.User,
			conf.Postgres.Password,
			conf.Postgres.URL,
			conf.Postgres.Database,
			conf.Postgres.SSLMode,
		)
		conn, err := sql.Open("postgres", url)
		if err != nil {
			logger.Error("error opening postgres connection", "error", err.Error())
			return
		}
		if err := conn.PingContext(ctx); err != nil {
			logger.Error("error verifying postgres connection", "error", err.Error())
			return
		}
		pgDb = conn

		migrator, err := migrate.New(conf.Postgres.MigrationsPath, url)
		if err != nil {
			logger.Error("error creating a new migrator", "error", err.Error())
			return
		}
		defer migrator.Close()

		if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			if errDirty, ok := err.(migrate.ErrDirty); ok {
				logger.Error(
					"migration is dirty, forcing rollback and retrying",
					"error", err.Error(),
				)
				migrator.Force(errDirty.Version - 1)
				if err := migrator.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
					logger.Error("migration error", "error", err.Error())
					return
				}
			} else {
				logger.Error("migration error", "error", err.Error())
				return
			}
		}
		logger.Info("database migrations completed successfully")
	}
}

func Postgres(ctx context.Context, logger logging.Logger) *sql.DB {
	pgInit.Do(connectPostgres(ctx, logger))
	return pgDb
}
