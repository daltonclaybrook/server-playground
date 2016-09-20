package controller

import (
	"fmt"
	"net/http"
)

const (
	Create  = iota
	Find    = iota
	FindOne = iota
	Update  = iota
	Delete  = iota
)

type CRUDRoute struct {
	Op      int
	Handler func(w http.ResponseWriter, r *http.Request)
}

func RoutesFromCRUD(model string, crud []CRUDRoute) []Route {

	m := make(map[string]*Route)
	for _, r := range crud {

		getRoute := func(pattern string) *Route {
			route := m[pattern]
			if route == nil {
				route = &Route{Path: pattern}
				m[pattern] = route
			}
			return route
		}

		switch r.Op {
		case Create:
			route := getRoute(fmt.Sprintf("/%v", model))
			route.Handlers = append(route.Handlers, Handler{"post", r.Handler})
		case Find:
			route := getRoute(fmt.Sprintf("/%v", model))
			route.Handlers = append(route.Handlers, Handler{"get", r.Handler})
		case FindOne:
			route := getRoute(fmt.Sprintf("/%v/{id:[0-9]+}", model))
			route.Handlers = append(route.Handlers, Handler{"get", r.Handler})
		case Update:
			route := getRoute(fmt.Sprintf("/%v/{id:[0-9]+}", model))
			route.Handlers = append(route.Handlers, Handler{"patch", r.Handler})
		case Delete:
			route := getRoute(fmt.Sprintf("/%v/{id:[0-9]+}", model))
			route.Handlers = append(route.Handlers, Handler{"delete", r.Handler})
		}
	}

	retRoutes := make([]Route, 0, len(m))
	for _, value := range m {
		retRoutes = append(retRoutes, *value)
	}

	return retRoutes
}
