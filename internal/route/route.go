package route

import (
	"fmt"
	"io"
	"net/http"
)

var GET = map[string]func(){
	"/ping": func() {
		fmt.Println("pong!")
	},
}

var POST = map[string]func(body io.ReadCloser){}

var PUT = map[string]func(*http.Request){}

var DELETE = map[string]func(*http.Request){}
