package app

import (
	"fmt"
	"github.com/levchenki/tea-app/internal/storage/postgres"
	"log/slog"
	"net/http"
	"os"

	"github.com/levchenki/tea-app/internal/config"
	"github.com/levchenki/tea-app/internal/routes"
)

func Run() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)
	logger.Info("Starting url-shortener", slog.String("env", cfg.Env))
	logger.Debug("Debug messages are enabled")

	err := postgres.RunMigrations(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	logger.Info("The migrations have been completed successfully")

	r := routes.SetupRouter(logger)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.HTTPServer.Port),
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		logger.Error("Failed to start the server")
	}
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case config.EnvDevelopment:
		log = slog.New(
			slog.NewTextHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelDebug,
				},
			),
		)
	case config.EnvProduction:
		log = slog.New(
			slog.NewJSONHandler(
				os.Stdout,
				&slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		)
	}
	return log
}
