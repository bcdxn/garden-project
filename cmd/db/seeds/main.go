package main

import (
	"context"
	"log"
	"os"

	"github.com/bcdxn/garden-project/internal/db/seeds"
)

func main() {
	dbURI := os.Getenv("DB_URI")

	if dbURI == "" {
		log.Fatal("missing required env var: DB_URI")
	}

	seeds.Roles(context.Background(), dbURI)
}
