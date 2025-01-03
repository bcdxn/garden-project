package main

import (
	"context"
	"net/http"
	"net/url"
	"os"

	rbac_app "github.com/bcdxn/garden-project/internal/app/rbac"
	"github.com/bcdxn/garden-project/internal/infrastructure/db/database"
	rbac_model "github.com/bcdxn/garden-project/internal/infrastructure/db/rbac"
	"github.com/bcdxn/garden-project/internal/infrastructure/http_middleware"
	"github.com/bcdxn/garden-project/internal/infrastructure/logger"
	"github.com/bcdxn/garden-project/internal/infrastructure/rest_api"
	"github.com/swaggest/swgui/v5emb"
)

func main() {
	ctx := context.Background()
	// todo: formalize config
	// instantiate application logger
	logger := logger.NewAppLogger([]any{http_middleware.RequestIDCtxKey})
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		logger.ErrorContext(ctx, "missing required env var DB_URI")
		os.Exit(1)
	}
	// Initialize DB connection
	db := database.Connect(ctx, logger, dbURI)
	// Instantiate Repository Implementations
	rbacRepository := &rbac_model.Model{DB: db}
	// Instantiate Services
	rbacService := rbac_app.NewService(rbacRepository)
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of
	//every operation from the generated code
	server := rest_api.NewServer(rbacService)
	router := http.NewServeMux()
	// serve open api yaml static file
	router.HandleFunc("/_docs/api/v1/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./internal/infrastructure/rest_api/openapi.yaml")
	})
	// add the open API Swagger Page
	router.Handle("/_docs/api/v1/", v5emb.New(
		"Garden Project",
		"/_docs/api/v1/openapi.yaml",
		"/_docs/api/v1/ui",
	))
	// Add middleware
	handler := rest_api.HandlerFromMux(server, router)
	handler = http_middleware.LogRequest(logger)(handler)
	handler = http_middleware.AddRequestID()(handler)
	addr := "http://0.0.0.0:8080"
	uri, err := url.Parse(addr)
	if err != nil {
		logger.ErrorContext(ctx, "invalid http server address", "err", err)
	}
	// instantiate http server
	s := &http.Server{
		Handler: handler,
		Addr:    uri.Host,
	}
	// Listen for HTTP requests
	logger.InfoContext(ctx, "http server is listening", "host", uri.Hostname(), "port", uri.Port())
	err = s.ListenAndServe()
	if err != nil {
		logger.ErrorContext(ctx, "error running http server", "err", err)
		os.Exit(1)
	}
}
