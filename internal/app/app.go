package app

import (
	"fmt"
	"github.com/levchenki/tea-app/internal/controller/http/routes"
	v1 "github.com/levchenki/tea-app/internal/controller/http/routes/v1"
	"github.com/levchenki/tea-app/internal/logger"
	"github.com/levchenki/tea-app/internal/storage/postgres"

	"github.com/levchenki/tea-app/internal/migrations"
	"github.com/levchenki/tea-app/internal/repository"

	"github.com/levchenki/tea-app/internal/service"
	"log/slog"
	"net/http"
	"os"

	"github.com/levchenki/tea-app/internal/config"
)

func Run() {
	cfg := config.MustLoad()

	lg := logger.SetupLogger(cfg.Env)
	lg.Info("Starting url-shortener", slog.String("env", cfg.Env))
	lg.Debug("Debug messages are enabled")

	migrations.Run(cfg, lg)

	db, err := postgres.New(cfg)
	if err != nil {
		//todo lg.Fatal()
		lg.Error("Failed to connect to database", err)
		os.Exit(1)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			lg.Error("Failed to close database connection", err)
			os.Exit(1)
		}
	}()

	teaService := service.NewTeaService(repository.NewTeaRepository(db))

	r := routes.SetupRouter(lg)
	r.Mount("/tea", v1.NewTeaRouter(lg, teaService))

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.HTTPServer.Port),
		Handler:      r,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		lg.Error("Failed to start the server")
	}
}
