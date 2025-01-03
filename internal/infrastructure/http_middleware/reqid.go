package http_middleware

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/segmentio/ksuid"
)

// custom key type to prevent collisions
type reqIDCtxKey string

const RequestIDCtxKey reqIDCtxKey = "reqid"

// RequestID returns a middleware that adds a unique request ID to each request context.
func AddRequestID() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), RequestIDCtxKey, slog.Attr{
				Key:   string(RequestIDCtxKey),
				Value: slog.StringValue("req_" + ksuid.New().String()),
			})
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
