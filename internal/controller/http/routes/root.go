package routes

import (
	lg "github.com/levchenki/tea-app/internal/controller/http/middleware/logger"
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(logger *slog.Logger) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(lg.New(logger))
	r.Use(middleware.Recoverer)

	return r
}
