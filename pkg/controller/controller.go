package controller

import "github.com/Rob9nn/gelin-game/pkg/router"

type Controller interface {
	Routes() []router.Route
}
