package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rob9nn/gelin-game/internal/player"
	"github.com/Rob9nn/gelin-game/pkg/controller"
)

type errorResponse struct {
	message string
}

type server struct {
	get    map[string]func(w http.ResponseWriter, r *http.Request)
	update map[string]func(w http.ResponseWriter, r *http.Request)
	delete map[string]func(w http.ResponseWriter, r *http.Request)
	post   map[string]func(w http.ResponseWriter, r *http.Request)
}

type Server interface {
	loadRoute(c controller.Controller)
}

var s = &server{
	get:    make(map[string]func(w http.ResponseWriter, r *http.Request)),
	update: make(map[string]func(w http.ResponseWriter, r *http.Request)),
	delete: make(map[string]func(w http.ResponseWriter, r *http.Request)),
	post:   make(map[string]func(w http.ResponseWriter, r *http.Request)),
}

// received set of controller ?
func Run() {
	log.Println("Listening on :8080")

	// load routes
	pc := player.PlayerController{}
	s.loadRoutes(pc)

	// set global function that handle everything
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// add routes to correct method for each controller
func (s *server) loadRoutes(c controller.Controller) {
	routes := c.Routes()
	for _, r := range routes {
		log.Printf("Load [%s] %s", r.Method_type, r.Path)
		s.getMethodMap(r.Method_type)[r.Path] = r.Handler
	}
}

func (s *server) getMethodMap(method string) map[string]func(w http.ResponseWriter, r *http.Request) {
	if method == http.MethodPost {
		return s.post
	} else if method == http.MethodDelete {
		return s.delete
	} else if method == http.MethodPut {
		return s.update
	}

	return s.get
}

func handler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	log.Printf("handling : [%s] %s", req.Method, req.RequestURI)
	handler, found := s.getMethodMap(req.Method)[req.RequestURI]
	if !found {
		routeNotFound(resp, req)
		return
	}

	// call handler
	handler(resp, req)
}

func routeNotFound(w http.ResponseWriter, r *http.Request) {
	e := errorResponse{
		message: "Route " + r.RequestURI + " not found.",
	}
	log.Println(e.message)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(e)
}
