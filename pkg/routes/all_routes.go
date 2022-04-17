package routes

import (
	s "chitest/pkg/services"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	r.Route("/ssh_keys", func(r chi.Router) {
		r.Get("/", s.ListSshKeys)
		r.Post("/", s.CreateArticle)
	})
}
