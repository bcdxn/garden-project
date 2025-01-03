// Logging middleware to log the incoming request and response status. It also contains the slog
// handler that enables logging context values (required by the logging middleware to log the
// request ID).

package http_middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

// RequestLogger returns a middleware that logs incoming request metadata.
func LogRequest(l *slog.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// capture the current time
			start := time.Now()
			// wrap the writer with a negroni writer so we can capture the response status code
			lrw := negroni.NewResponseWriter(w)
			h.ServeHTTP(lrw, r)
			// calculate duration of API invocation since 'start'
			duration := float64(time.Since(start)) / float64(time.Millisecond)
			// log the request
			l.InfoContext(
				r.Context(),
				"request",
				"status", lrw.Status(),
				"method", r.Method,
				"path", r.URL.Path,
				"duration_ms", duration,
			)
		})
	}
}
