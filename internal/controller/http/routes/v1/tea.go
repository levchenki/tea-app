package v1

import (
	"fmt"
	"github.com/levchenki/tea-app/internal/entity"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type TeaService interface {
	Get() ([]entity.Tea, error)
	GetByTeaId(id int) (entity.Tea, error)
	GetByCategoryId(id int) ([]entity.Tea, error)
}

func NewTeaRouter(logger *slog.Logger, s TeaService) chi.Router {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			teas, err := s.Get()
			if err != nil {
				logger.Error("Failed to get all teas", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			logger.Info("Successfully got all teas", teas)
			_, err = w.Write([]byte(fmt.Sprintf("%v", teas)))
			if err != nil {
				logger.Error("Failed to write response", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		})

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			idInt, err := strconv.Atoi(id)
			if err != nil {
				logger.Error("Failed to convert id to int", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			tea, err := s.GetByTeaId(idInt)
			if err != nil {
				logger.Error("Failed to get tea by id", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Print(tea)
			logger.Info("Successfully got tea by id")
			w.Write([]byte(fmt.Sprintf("%v", tea)))
		})

		r.Get("/category/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chi.URLParam(r, "id")
			idInt, err := strconv.Atoi(id)
			if err != nil {
				logger.Error("Failed to convert id to int", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			teas, err := s.GetByCategoryId(idInt)
			if err != nil {
				logger.Error("Failed to get all teas by category id", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Print(teas)
			logger.Info("Successfully got all teas by category id")
			w.Write([]byte(fmt.Sprintf("%v", teas)))
		})
	})

	//router.Group(func(r chi.Router) {
	//	r.Use(admin.AdminOnly)
	//
	//	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
	//		w.Write([]byte("teas: creating a tea"))
	//	})
	//
	//	r.Put("/", func(w http.ResponseWriter, r *http.Request) {
	//		w.Write([]byte("teas: updating a tea"))
	//	})
	//
	//	r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
	//		w.Write([]byte("teas: deleting a tea"))
	//	})
	//})

	return router
}
