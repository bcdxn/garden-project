package main

import (
	"log"
	"net/http"

	rbac_app "github.com/bcdxn/garden-project/internal/app/rbac"
	rbac_model "github.com/bcdxn/garden-project/internal/infrastructure/db/rbac"
	"github.com/bcdxn/garden-project/internal/infrastructure/rest_api"
)

func main() {
	// Instantiate Repository Implementations
	rbacRepository := &rbac_model.Model{}
	// Instantiate Services
	rbacService := rbac_app.NewService(rbacRepository)
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := rest_api.NewServer(rbacService)

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := rest_api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
