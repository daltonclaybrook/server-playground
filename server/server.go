package server

import (
	"fmt"
	"github.com/daltonclaybrook/web-app/controller"
	"net/http"
	"strings"
	"time"
)

// WebServer is used to create and start a server.
type WebServer struct {
	server *http.Server
}

// RegisterController registers a request handler with the WebServer.
func (ws *WebServer) RegisterController(c controller.Controller) {

	routes := c.Routes()
	for i := 0; i < len(routes); i++ {
		route := routes[i]
		http.HandleFunc(route.Path, func(writer http.ResponseWriter, request *http.Request) {
			method := strings.ToLower(request.Method)
			for j := 0; j < len(route.Handlers); j++ {
				handler := route.Handlers[j]
				if strings.ToLower(handler.Method) == method {
					handler.Handler(writer, request)
					return
				}
			}

			sendUnhandled(writer, request)
		})
	}
}

// Start starts the server.
func (ws *WebServer) Start() {
	ws.setupServer()
	ws.server.ListenAndServe()
}

/*
Private
*/

func (ws *WebServer) setupServer() {
	ws.server = &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func sendUnhandled(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	fmt.Fprintf(w, "Method \"%v\" is not supported by this route.", r.Method)
}
