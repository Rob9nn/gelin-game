package main

import (
	"github.com/Rob9nn/gelin-game/internal/db"
	"github.com/Rob9nn/gelin-game/internal/migration"
)

func main() {
	db.NewConnectionPool(1)
	migration.Migrate()
}
