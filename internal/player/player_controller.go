package player

import (
	"log"
	"net/http"

	"github.com/Rob9nn/gelin-game/pkg/router"
)

type PlayerController struct{}

func (PlayerController) create(w http.ResponseWriter, r *http.Request) {
	log.Printf("create new player yooo")
}

func (PlayerController) update(w http.ResponseWriter, r *http.Request) {
}

func (PlayerController) delete(w http.ResponseWriter, r *http.Request) {
}

func (PlayerController) read(w http.ResponseWriter, r *http.Request) {
}

func (PlayerController) Routes() []router.Route {
	return []router.Route{
		{Method_type: http.MethodPost, Path: "/v1/player/create", Handler: PlayerController{}.create},
		{Method_type: http.MethodPut, Path: "/v1/player/update", Handler: PlayerController{}.update},
		{Method_type: http.MethodDelete, Path: "/v1/player/delete", Handler: PlayerController{}.delete},
		{Method_type: http.MethodGet, Path: "/v1/player/read", Handler: PlayerController{}.read},
	}
}
