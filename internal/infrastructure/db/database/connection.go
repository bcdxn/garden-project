package database

import (
	"context"
	"database/sql"
	"log/slog"
	"net/url"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(ctx context.Context, logger *slog.Logger, dbURI string) *sql.DB {
	url, err := url.Parse(dbURI)
	if err != nil {
		logger.ErrorContext(ctx, "invalid DB URI", "err", err)
	}
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		logger.ErrorContext(ctx, "unable to connect to database", "err", err)
		os.Exit(1)
	}
	logger.InfoContext(ctx, "successfully connected to database", "host", url.Hostname(), "port", url.Port())
	return db
}
