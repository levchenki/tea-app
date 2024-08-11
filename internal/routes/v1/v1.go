package v1

import "github.com/go-chi/chi/v5"

func SetupV1Router() chi.Router {
	r := chi.NewRouter()

	r.Mount("/teas", teaRouter())
	r.Mount("/categories", categoryRouter())
	r.Mount("/users", userRouter())

	return r
}
