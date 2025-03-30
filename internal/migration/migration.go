package migration

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Rob9nn/gelin-game/internal/db"
	"github.com/Rob9nn/gelin-game/resources"
)

var migrationTableName = "migration_history"

func Migrate() {
	// get a DB connection and start a transaction
	conn, err := db.GetConnection()
	if err != nil {
		log.Panic(err)
	}
	defer db.CloseConnection()

	tx, err := conn.Begin()
	if err != nil {
		log.Panic(err)
	}

	createMigrationTableIfNotExists(tx)
	tx.Commit()

	tx, err = conn.Begin()
	if err != nil {
		log.Panic(err)
	}
	if err := loadMigration(tx); err != nil {
		tx.Rollback()
		log.Print(err)
	}

	tx.Commit()
}

func loadMigration(conn *sql.Tx) error {
	entries, err := resources.GetSqlFiles()
	if err != nil {
		log.Panic(err)
	}

	h := sha256.New() // used to calculate checksum
	for _, entry := range entries {
		if entry.Name()[0] == 'V' {
			log.Println("Versioned file")
			fileNameSplit := strings.Split(entry.Name(), "__")
			if len(fileNameSplit) == 1 {
				log.Printf("invalid version format detected : %s", entry.Name())
				continue
			}

			isNewVersion, err := isNewVersion(fileNameSplit[0], conn)
			if err != nil {
				return err
			}
			if isNewVersion {
				fileContent, err := resources.ReadSqlFile(entry.Name())
				if err != nil {
					return fmt.Errorf("error while reading %s", err)
				}
				_, err = conn.Exec(string(fileContent[:]))
				if err != nil {
					return fmt.Errorf("error while executing file content %s", err)
				}
				err = addToMigrationHistory(fileNameSplit, h.Sum(fileContent), conn)
				if err != nil {
					return fmt.Errorf("error while adding to migration history %s", err)
				}
			}
		} else if entry.Name()[0] == 'R' {
			// execute and add to migration_history
			log.Println("Repeatable file")
		}
	}

	return nil
}

func addToMigrationHistory(fileNameSplit []string, checksum []byte, conn *sql.Tx) error {
	query := fmt.Sprintf(`insert into %s 
		(version, description, checksum) 
		values ($1, $2, $3)`, migrationTableName)
	description := strings.ReplaceAll(fileNameSplit[1], "_", " ")
	description = strings.ReplaceAll(description, ".sql", "")

	_, err := conn.Query(query, fileNameSplit[0], description, checksum)
	return err
}

func isNewVersion(version string, conn *sql.Tx) (bool, error) {
	var count int
	query := fmt.Sprintf("select count(1) from %s where version = $1", migrationTableName)
	err := conn.QueryRow(query, version).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("isNewVerison: %s", err)
	}

	return count == 0, nil
}

func createMigrationTableIfNotExists(conn *sql.Tx) {
	var count int8
	query := "select count(1) from pg_catalog.pg_tables where tablename = '%s'"
	err := conn.QueryRow(fmt.Sprintf(query, migrationTableName)).Scan(&count)
	if err != nil {
		conn.Rollback()
		log.Fatal(err)
	}

	if count == 0 {
		//	--> V1.01__my_description.sql
		//	--> R__my_description.sql = R => repeatable
		//	--> U__my_description.sql = undo
		log.Printf("create %s table.", migrationTableName)
		query := `CREATE TABLE %s (
			migration_history_id serial primary key,
			version text not null,
			description text not null,
			checksum bytea not null,
			migrated_at timestamp not null default now()
		)`
		_, err := conn.Exec(fmt.Sprintf(query, migrationTableName))
		if err != nil {
			conn.Rollback()
			log.Fatal(err)
		}
	}
}
