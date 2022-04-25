package services

import (
	"chitest/pkg/ctxval"
	"net/http"

	"github.com/rs/zerolog"
)

func ContextLogger(r *http.Request) zerolog.Logger {
	return r.Context().Value(ctxval.LoggerCtxKey).(zerolog.Logger)
}
