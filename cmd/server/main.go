package main

import (
	"log"

	"github.com/Rob9nn/gelin-game/internal/db"
	"github.com/Rob9nn/gelin-game/internal/migration"
	"github.com/Rob9nn/gelin-game/internal/server"
)

func main() {
	err := db.NewConnectionPool(10)
	if err != nil {
		log.Panic(err)
	}
	migration.Migrate()
	server.Run()
}
