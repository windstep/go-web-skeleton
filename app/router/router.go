package router

import (
	"github.com/gorilla/mux"
	"github.com/windstep/go-web-skeleton/app/router/routes"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddleware(r)
}
