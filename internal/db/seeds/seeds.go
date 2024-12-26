package seeds

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/segmentio/ksuid"
)

const (
	insertRoleQuery = `
INSERT INTO rbac_role(id, name)
VALUES($1, $2)
`
	querySeedsVersion = `
SELECT version, dirty FROM seeds
`

	insertSeedsVersion = `
INSERT INTO seeds(version, dirty)
VALUES($1, $2)
`
)

const rolesSchemaSeedsVersion = 1

func Roles(ctx context.Context, dbURI string) error {
	// Open a database/sql connection using pgx
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatalf("error connecting to database - %s", err.Error())
	}

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var sr seedsRecord
	rows := tx.QueryRowContext(ctx, querySeedsVersion)
	err = rows.Scan(&sr)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	// the seeding for roles has already run
	if sr.Version >= rolesSchemaSeedsVersion {
		return nil
	}
	// reset err variable so that tx can be committed
	err = nil

	roles := []struct {
		id   string
		name string
	}{
		{
			id:   "rol_" + ksuid.New().String(),
			name: "user",
		},
		{
			id:   "rol_" + ksuid.New().String(),
			name: "verified_user",
		},
		{
			id:   "rol_" + ksuid.New().String(),
			name: "admin",
		},
	}

	for _, role := range roles {
		_, err := tx.QueryContext(ctx, insertRoleQuery, role.id, role.name)
		if err != nil {
			log.Fatalf("error inserting role - %s", err.Error())
		}
	}

	_, err = tx.QueryContext(ctx, insertSeedsVersion, 1, false)

	return nil
}

type seedsRecord struct {
	Version int  `sql:"version"`
	Dirty   bool `sql:"dirty"`
}
