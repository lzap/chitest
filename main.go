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
	// initialize stdout logging and AWS clients first
	log.Logger = logging.InitializeStdout()
	aws.Initialize()

	// initialize cloudwatch using the AWS clients
	logger, clsFunc, err := logging.InitializeCloudwatch(log.Logger)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer clsFunc()
	log.Logger = logger

	// initialize the rest
	db.Initialize()

	r := chi.NewRouter()
	r.Use(m.RequestID)
	r.Use(m.RequestNum)
	r.Use(middleware.URLFormat)
	r.Use(m.LoggerMiddleware(&log.Logger))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	routes.SetupRoutes(r)

	log.Info().Msg("New instance started ")
	http.ListenAndServe(":3000", r)
}
