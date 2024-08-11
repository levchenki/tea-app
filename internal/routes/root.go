package routes

import (
	"log/slog"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	lg "github.com/levchenki/tea-app/internal/http-server/middleware/logger"
	v1 "github.com/levchenki/tea-app/internal/routes/v1"
)

func SetupRouter(logger *slog.Logger) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(lg.New(logger))
	r.Use(middleware.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Mount("/v1", v1.SetupV1Router())
	})

	return r
}
