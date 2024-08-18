package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/levchenki/tea-app/internal/config"
	"github.com/pkg/errors"
)

func New(cfg *config.Config) (*sqlx.DB, error) {
	connectionString := GetDatabaseURL(cfg)

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

func GetDatabaseURL(cfg *config.Config) string {
	storage := cfg.Storage
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		storage.Username,
		storage.Password,
		storage.Host,
		5432,
		storage.Name,
	)
}
