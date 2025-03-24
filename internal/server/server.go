package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Rob9nn/gelin-game/internal/route"
)

type errorResponse struct {
	message string
}

func Run() {
	writeHeader()
	log.Println("Listening on :8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func writeHeader() {
	dir, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	fileName := "/internal/server/server-header.txt"
	data, err := os.ReadFile(dir + fileName)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%s v.0.0.1\n", data)
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
