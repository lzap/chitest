package middleware

import (
	"chitest/pkg/ctxval"
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

var panicStatus = http.StatusInternalServerError

func LoggerMiddleware(rootLogger *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			bytesIn, _ := strconv.Atoi(r.Header.Get("Content-Length"))
			rid := ctxval.GetValue(r.Context(), ctxval.RequestIdCtxKey)
			logger := rootLogger.With().
				Timestamp().
				Str("rid", rid).
				Str("remote_ip", r.RemoteAddr).
				Str("url", r.URL.Path).
				Str("method", r.Method).
				Int("bytes_in", bytesIn).
				Logger()
			t1 := time.Now()

			defer func() {
				duration := time.Since(t1)
				log := logger.With().
					Dur("latency_ms", duration).
					Int("bytes_out", ww.BytesWritten()).
					Logger()

				if rec := recover(); rec != nil {
					log.Error().
						Bool("panic", true).
						Int("status", panicStatus).
						Interface("recover_info", rec).
						Bytes("debug_stack", debug.Stack()).
						Msg("Unhandled panic")
					http.Error(ww, http.StatusText(panicStatus), panicStatus)
				}

				log.Info().
					Int("status", ww.Status()).
					Msg(fmt.Sprintf("Completed %s request %s in %d ms with %d",
						r.Method, r.URL.Path, duration.Milliseconds(), ww.Status()))
			}()

			ctx := context.WithValue(r.Context(), ctxval.LoggerCtxKey, logger)
			next.ServeHTTP(ww, r.WithContext(ctx))
		}
		return http.HandlerFunc(fn)
	}
}