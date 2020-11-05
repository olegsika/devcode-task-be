package rest

import (
	"github.com/gorilla/mux"
)

type Controller struct {
	Prefix       string
	Routes       []*Route
}

func (c *Controller) InitRoutes(r *mux.Router) {
	var pR = r

	if c.Prefix != "" {
		pR = r.PathPrefix("/" + c.Prefix).Subrouter()
	}

	for _, route := range c.Routes {
		pR.HandleFunc(route.Path, route.Action).Methods(route.Method, "OPTIONS")
	}
}
