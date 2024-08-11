package app

import (
	"fmt"
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

	r := routes.SetupRouter(logger)
	port := cfg.HTTPServer.Port

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
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
