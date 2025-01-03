package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(dbURI string) *sql.DB {
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		panic(err)
	}
	return db
}
