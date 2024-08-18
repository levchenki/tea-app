package migrations

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/levchenki/tea-app/internal/config"
	"github.com/levchenki/tea-app/internal/storage/postgres"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

func Run(cfg *config.Config, logger *slog.Logger) {
	databaseURL := postgres.GetDatabaseURL(cfg)

	m, err := migrate.New("file://database/migrations", databaseURL)

	if err != nil {
		logger.Error("Migration Error", err)
		os.Exit(1)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Error("Failed to run migration", err)
		os.Exit(1)
	}

	defer func() {
		sourceError, databaseError := m.Close()
		if sourceError != nil {
			logger.Error("Failed to close migration source", sourceError)
			os.Exit(1)
		}
		if databaseError != nil {
			logger.Error("Failed to close migration database", databaseError)
			os.Exit(1)
		}
	}()

	logger.Info("The migrations have been completed successfully")

}
