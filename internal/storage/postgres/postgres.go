package postgres

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/levchenki/tea-app/internal/config"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func New(cfg *config.Config) (*sqlx.DB, error) {
	connectionString := getConnectionString(cfg)

	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open connection to postgres")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to ping postgres")
	}

	return db, nil
}

func RunMigrations(cfg *config.Config) error {

	db, err := New(cfg)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})

	if err != nil {
		return errors.Wrap(err, "Failed to create database driver")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrations",
		"postgres", driver)

	if err != nil {
		return errors.Wrap(err, "Failed to create migration instance")
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return errors.Wrap(err, "Failed to run migrations")
	}
	return nil
}

func getConnectionString(cfg *config.Config) string {
	storage := cfg.Storage
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		storage.Username,
		storage.Password,
		storage.Host,
		5432,
		storage.Name,
	)
}
