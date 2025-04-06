package player

import (
	"net/http"

	"github.com/Rob9nn/gelin-game/pkg/server"
)

type PlayerController struct{}

func (PlayerController) create(w http.ResponseWriter, r *http.Request) {
}

func (PlayerController) update(w http.ResponseWriter, r *http.Request) {
}

func (PlayerController) delete(w http.ResponseWriter, r *http.Request) {
}

func (PlayerController) read(w http.ResponseWriter, r *http.Request) {
}

func (PlayerController) Routes() []server.Route {
	return []server.Route{
		{Method_type: "POST", Path: "v1/player/create", Handler: PlayerController{}.create},
		{Method_type: "PUT", Path: "v1/player/update", Handler: PlayerController{}.update},
		{Method_type: "DELETE", Path: "v1/player/delete", Handler: PlayerController{}.delete},
		{Method_type: "GET", Path: "v1/player/read", Handler: PlayerController{}.read},
	}
}
