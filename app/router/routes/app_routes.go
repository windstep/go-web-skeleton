package routes

import (
	"github.com/windstep/go-web-skeleton/app/api_controllers"
	"github.com/windstep/go-web-skeleton/app/controllers"
	"github.com/windstep/go-web-skeleton/app/middleware"
)

var appRoutes = []Route{
	{
		Url:             "/app",
		Method:          "GET",
		Controller:      api_controllers.HealthCheck,
		MiddlewareChain: middleware.ApiMiddlewareChain,
	},
	{
		Url:             "/",
		Method:          "GET",
		Controller:      controllers.HealthCheck,
		MiddlewareChain: middleware.WebMiddlewareChain,
	},
}
