package logger

import (
	"context"
	"log/slog"
	"os"
)

func NewAppLogger(keysToLog []any) *slog.Logger {
	encoderHandler := slog.NewJSONHandler(os.Stdout, nil)
	contextHandler := slogContextHandler{
		Handler: encoderHandler,
		Keys:    keysToLog,
	}

	return slog.New(contextHandler)
}

type slogContextHandler struct {
	slog.Handler
	Keys []any
}

func (h slogContextHandler) Handle(ctx context.Context, r slog.Record) error {
	r.AddAttrs(h.observe(ctx)...)
	return h.Handler.Handle(ctx, r)
}

func (h slogContextHandler) observe(ctx context.Context) (as []slog.Attr) {
	for _, k := range h.Keys {
		a, ok := ctx.Value(k).(slog.Attr)
		if !ok {
			continue
		}
		a.Value = a.Value.Resolve()
		as = append(as, a)
	}
	return
}
