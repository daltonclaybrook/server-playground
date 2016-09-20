package controller

import (
	"net/http"
)

// Controller handles routes.
type Controller interface {
	Routes() []Route
}

// Route describes an endpoint.
type Route struct {
	Path     string
	Handlers []Handler
}

// Handler describes functions mapped to http methods.
type Handler struct {
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}
