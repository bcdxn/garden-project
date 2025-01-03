package main

import (
	"log"
	"net/http"
	"os"

	rbac_app "github.com/bcdxn/garden-project/internal/app/rbac"
	"github.com/bcdxn/garden-project/internal/infrastructure/db/database"
	rbac_model "github.com/bcdxn/garden-project/internal/infrastructure/db/rbac"
	"github.com/bcdxn/garden-project/internal/infrastructure/rest_api"
	"github.com/swaggest/swgui/v5emb"
)

func main() {
	// todo: formalize config
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("missing required env var DB_URI")
	}
	// Initialize DB connection
	db := database.Connect(dbURI)
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
	// instantiate http server
	s := &http.Server{
		Handler: rest_api.HandlerFromMux(server, router),
		Addr:    "0.0.0.0:8080",
	}
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}