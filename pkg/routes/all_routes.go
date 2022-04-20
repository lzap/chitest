package routes

import (
	"chitest/pkg/middleware"
	s "chitest/pkg/services"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux) {
	r.Route("/ssh_keys", func(r chi.Router) {
		r.Get("/", s.ListSshKeys)
		r.Post("/", s.CreateSShKey)
		r.Route("/{ID}", func(r chi.Router) {
			r.Use(middleware.SshKeyCtx)
			r.Get("/", s.GetSshKey)
			r.Put("/", s.UpdateSshKey)
			r.Delete("/", s.DeleteSshKey)
			r.Route("/resources", func(r chi.Router) {
				r.Get("/", s.ListSshKeyResources)
				r.Post("/", s.CreateSshKeyResource)
				r.Route("/{RID}", func(r chi.Router) {
					r.Use(middleware.SshKeyResourceCtx)
					r.Delete("/", s.DeleteSshKeyResource)
				})
			})
		})
	})
}
