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
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	pgInit sync.Once
	pgDb   *sql.DB
)

func newPostgresConn(ctx context.Context) (*sql.DB, error) {
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
		return nil, err
	}
	if err := conn.PingContext(ctx); err != nil {
		return nil, err
	}
	return conn, err
}

func Postgres(ctx context.Context, logger logging.Logger) *sql.DB {
	pgInit.Do(func() {
		conn, err := newPostgresConn(ctx)
		if err != nil {
			logger.Error("error while connecting to postgres")
			panic(err)
		}
		migrationsPath := config.Get().Postgres.MigrationsPath
		if err := migrateUp(conn, migrationsPath, logger); err != nil {
			logger.Error("error while applying migrations")
			panic(err)
		}
	})
	return pgDb
}

func migrateUp(db *sql.DB, migrations string, logger logging.Logger) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	migrator, err := migrate.NewWithDatabaseInstance(
		migrations,
		"postgres",
		driver,
	)
	if err != nil {
		logger.Error("error creating a new migrator")
		return err
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
				logger.Error("migration error")
				return err
			}
		} else {
			logger.Error("migration error")
			return err
		}
	}
	logger.Info("database migrations completed successfully")
	return nil
}
