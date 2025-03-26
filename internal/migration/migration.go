package migration

import (
	"database/sql"
	"log"
	"os"

	"github.com/Rob9nn/gelin-game/internal/db"
)

func Migrate() {
	// get a DB connection and start a transaction
	conn, err := db.GetConnection()
	if err != nil {
		log.Panic(err)
	}

	createMigrationTableIfNotExists(conn)

	// let this here to be able to configure it
	dirPath :=
		loadMigration(dirPath)
}

func loadMigration(dirPath string) {
	// go through dir and execute files.
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Panic(err)
	}

	for _, entry := range entries {
		log.Printf("[debug] " + entry.Name())
	}
}

func createMigrationTableIfNotExists(conn *sql.DB) {
	var count int8
	err := conn.QueryRow("select count(1) from pg_catalog.pg_tables where tablename = 'migration_history'").Scan(&count)
	if err != nil {
		log.Panic(err)
	}

	if count == 0 {
		//	--> V1.01__my_description.sql
		//	--> R__my_description.sql = R => repeatable
		//	--> U__my_description.sql = undo
		log.Printf("create migration history table.")
		_, err := conn.Exec(`CREATE TABLE migration_history (
			migration_history_id serial primary key,
			version text not null,
			description text not null,
			checksum integer not null,
			migrated_at timestamp not null default now()
		)`)
		if err != nil {
			log.Panic(err)
		}
	}
}
