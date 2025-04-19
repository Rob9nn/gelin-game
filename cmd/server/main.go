package main

import (
	"log"

	"github.com/Rob9nn/gelin-game/pkg/db"
	"github.com/Rob9nn/gelin-game/pkg/migration"
	"github.com/Rob9nn/gelin-game/pkg/server"
)

func main() {
	err := db.NewConnectionPool(10)
	if err != nil {
		log.Panic(err)
	}
	migration.Migrate()
	server.Run()
}
