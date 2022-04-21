package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

var panicStatus = http.StatusInternalServerError

func LoggerMiddleware(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			defer func() {
				duration := time.Since(t1)
				bytes_in, _ := strconv.Atoi(r.Header.Get("Content-Length"))
				log := logger.With().
					Timestamp().
					Dur("latency_ms", duration).
					Str("remote_ip", r.RemoteAddr).
					Str("url", r.URL.Path).
					Str("method", r.Method).
					Int("bytes_in", bytes_in).
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

			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
