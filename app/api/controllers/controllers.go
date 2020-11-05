package controllers

import (
	"devcodeIdentity/app/api/controllers/actors"
	"devcodeIdentity/app/api/controllers/movies"
	"github.com/gorilla/mux"
	"sync"
)

var controllers []func(r *mux.Router)

var once sync.Once

func GetRoutes(r *mux.Router) {
	once.Do(func() {
		controllers = getControllers()
	})

	for _, controller := range controllers {
		controller(r)
	}
}

func getControllers() []func(r *mux.Router) {
	return []func(r *mux.Router){
		movies.Controller.InitRoutes,
		actors.Controller.InitRoutes,
	}
}
