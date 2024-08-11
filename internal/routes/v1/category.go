package v1

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/levchenki/tea-app/internal/http-server/middleware/admin"
)

func categoryRouter() chi.Router {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("category: all categories"))
		})

		r.Get("/{categoryId}", func(w http.ResponseWriter, r *http.Request) {
			categoryId := chi.URLParam(r, "categoryId")
			w.Write([]byte(fmt.Sprintf("category: %v", categoryId)))
		})
	})

	router.Group(func(r chi.Router) {
		r.Use(admin.AdminOnly)

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("category: creating a category"))
		})

		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("category: updating a category"))
		})

		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("category: deleting a category"))
		})
	})

	return router
}
