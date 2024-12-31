package main

import (
	"context"
	"log"
	"os"

	"github.com/bcdxn/garden-project/internal/infrastructure/db/seeds"
)

func main() {
	dbURI := os.Getenv("DB_URI")

	if dbURI == "" {
		log.Fatal("missing required env var: DB_URI")
	}

	seeds.Run(context.Background(), dbURI, seeds.Seed_000001_rbac)
}
