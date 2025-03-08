package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Rob9nn/gelin-game/internal/route"
)

type errorResponse struct {
	message string
}

func Run() {
	log.Println("Listening on :8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	if req.Method == "GET" {
		f, ok := route.GET[req.RequestURI]
		if !ok {
			routeNotFound(resp, req)
			return
		}
		f()
	} else if req.Method == "POST" {
		f, ok := route.POST[req.RequestURI]
		if !ok {
			routeNotFound(resp, req)
			return
		}
		f(req.Body)
	}
}

func routeNotFound(w http.ResponseWriter, r *http.Request) {
	e := errorResponse{
		message: "Route " + r.RequestURI + " not found.",
	}
	log.Println(e.message)
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(e)
}
