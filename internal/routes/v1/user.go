package v1

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/levchenki/tea-app/internal/http-server/middleware/admin"
)

func userRouter() chi.Router {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(admin.AdminOnly)

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("users: all users"))
		})

		r.Get("/{userId}", func(w http.ResponseWriter, r *http.Request) {
			userId := chi.URLParam(r, "userId")
			w.Write([]byte(fmt.Sprintf("users: user with id %v", userId)))
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("users: creating a user"))
		})

		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("users: deleting a user"))
		})
	})

	return router
}
