package logger

import (
	"github.com/levchenki/tea-app/internal/config"
	"log/slog"
	"os"
)

func SetupLogger(env string) *slog.Logger {

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
