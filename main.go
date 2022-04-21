package main

import (
	"chitest/pkg/clouds/aws"
	"chitest/pkg/db"
	"chitest/pkg/logging"
	m "chitest/pkg/middleware"
	"chitest/pkg/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func main() {
	logging.Initialize()
	db.Initialize()
	aws.Initialize()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(m.LoggerMiddleware(&log.Logger))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	routes.SetupRoutes(r)

	http.ListenAndServe(":3000", r)
}
