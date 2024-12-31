package seeds

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	querySeedsVersion  = "SELECT version FROM seeds"
	deleteSeedsVersion = "DELETE FROM seeds"
	insertSeedsVersion = "INSERT INTO seeds(version, dirty) VALUES($1, $2)"
)

func Run(ctx context.Context, dbURI string, s seed) error {
	// 0pen a database/sql connection using pgx
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatalf("error connecting to database - %s", err.Error())
	}
	// start a transaction
	tx, err := conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	// commit or rollback the transaction depending on the error status on function exit
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	// load the seeds version record to check where we are in the seeding steps
	var currentVersion int
	rows := tx.QueryRowContext(ctx, querySeedsVersion)
	err = rows.Scan(&currentVersion)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	// check if the seeding for this particular step has already run
	if currentVersion >= s.Version {
		return nil
	}
	// reset err variable so that tx can be committed
	err = nil
	// iterate over each step in the seed set
	for _, step := range s.Steps {
		// iterate over the data set and insert into the database
		for _, data := range step.Data {
			_, err := tx.QueryContext(ctx, step.SQL, data...)
			if err != nil {
				log.Fatalf("error seeding record - %s", err.Error())
			}
		}
	}
	// delete the old seeds version
	_, err = tx.QueryContext(ctx, deleteSeedsVersion)
	// insert the new seeds version
	_, err = tx.QueryContext(ctx, insertSeedsVersion, s.Version, false)

	return nil
}

/* Internal Types
------------------------------------------------------------------------------------------------- */

type seed struct {
	Version int
	Steps   []seedStep
}

type seedStep struct {
	SQL  string
	Data [][]any
}

type seedData struct {
	SQL  string
	Args []any
}

type seedsVersion struct {
	Version int  `sql:"version"`
	Dirty   bool `sql:"dirty"`
}
