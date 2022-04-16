package main

import (
	"chitest/pkg/db"
	"chitest/pkg/services"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	db.InitDatabase()
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/ssh_keys", func(r chi.Router) {
		r.Get("/", services.ListSshKeys)
		//r.Post("/", ListSshKeys)
	})

	http.ListenAndServe(":3000", r)
}
