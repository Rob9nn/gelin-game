package main

import (
	"github.com/Rob9nn/gelin-game/pkg/db"
	"github.com/Rob9nn/gelin-game/pkg/migration"
)

func main() {
	db.NewConnectionPool(1)
	migration.Migrate()
}
