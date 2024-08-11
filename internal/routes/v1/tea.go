package v1

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/levchenki/tea-app/internal/http-server/middleware/admin"
)

func teaRouter() chi.Router {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("teas: all teas"))
		})

		r.Get("/{teaId}", func(w http.ResponseWriter, r *http.Request) {
			teaId := chi.URLParam(r, "teaId")
			w.Write([]byte(fmt.Sprintf("teas: tea with id %v", teaId)))
		})
	})

	router.Group(func(r chi.Router) {
		r.Use(admin.AdminOnly)

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("teas: creating a tea"))
		})

		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("teas: updating a tea"))
		})

		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("teas: deleting a tea"))
		})
	})

	return router
}
