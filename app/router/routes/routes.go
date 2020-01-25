package routes

import (
	"github.com/windstep/go-web-skeleton/app/middleware"
	"github.com/windstep/go-web-skeleton/app/middleware/chain"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Url             string
	Method          string
	Controller      func(http.ResponseWriter, *http.Request)
	MiddlewareChain chain.Chain
}

func load() []Route {
	routes := appRoutes
	return routes
}

func SetupRoutesWithMiddleware(r *mux.Router) *mux.Router {
	for _, route := range load() {
		r.Handle(route.Url, route.MiddlewareChain.ThenFunc(route.Controller)).Methods(route.Method)
	}
	r.PathPrefix("/static/").Handler(
		middleware.LoggerMiddleware(http.StripPrefix("/static/", http.FileServer(http.Dir("/static"))).ServeHTTP)).Methods("GET")

	return r
}
