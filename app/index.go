package main

import (
	apiControllers "devcodeIdentity/app/api/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Part struct {
	Prefix      string
	Router      *mux.Router
	Controllers func(r *mux.Router)
}

func main() {

	r := mux.NewRouter()

	var api = Part{
		Prefix:      "/api",
		Controllers: apiControllers.GetRoutes,
	}

	api.Router = r.PathPrefix(api.Prefix).Subrouter()
	api.Controllers(api.Router)

	http.ListenAndServe(":3001", handlers.CompressHandler(handlers.CombinedLoggingHandler(os.Stdout, r)))
}
