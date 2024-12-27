package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	directionUp   = "up"
	directionDown = "down"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("direction is required, e.g. go run ./cmd/migrations/main.go up")
	}

	direction := args[0]

	if direction != directionUp && direction != directionDown {
		log.Fatalf("direction must be one of 'up' or 'down'")
	}

	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("missing required env var DB_URI")
	}

	m, err := migrate.New(
		"file://./internal/db/migrations",
		dbURI)
	if err != nil {
		log.Fatal("error creating migration instance", err)
	}

	if direction == directionUp {
		err = m.Up()
	} else {
		err = m.Down()
	}

	if err == migrate.ErrNoChange {
		log.Println("Migration already at latest; no changes")
	} else if err != nil {
		log.Fatal("error running migration: ", err)
	} else {
		log.Println("Migration ran successfully")
	}
	version, dirty, _ := m.Version()
	log.Println("Migration result:")
	log.Printf("\tVersion: %d\n", version)
	log.Printf("\tIs Dirty: %v\n", dirty)
}
