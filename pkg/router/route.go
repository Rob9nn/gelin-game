package router

import "net/http"

type Route struct {
	Method_type string
	Path        string
	Handler     http.HandlerFunc
}
