package migration

import (
	"log"

	"github.com/Rob9nn/gelin-game/internal/db"
)

func Migrate() {
	// get a DB connection and start a transaction
	conn, err := db.GetConnection()
	if err != nil {
		log.Panic(err)
	}

	var count int8
	err = conn.QueryRow("select count(1) from pg_catalog.pg_tables where tablename = 'migration_history'").Scan(&count)
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

	// check dir where migrate are and load
}
