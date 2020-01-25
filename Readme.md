# github.com/windstep/go-web-skeleton

Simple and clean web template for golang web programming. 
This boilerplate / skeleton takes care of basic logging, routing (middleware + route), configuration, and basic database 
connection and dockerizing an app. 

## how to use

Simply clone this repo, rename package name into your's name and start docker (just replace github.com/windstep/go-web-skeleton in go.mod file and in all project files).

## use with docker

`docker-compose up --build go` or `docker build . -t project:latest && docker run project:latest`

There is a database included into docker-compose file, so you can generally use it.

## what involved?

`main.go` is your entrance point. It just creates logger and config and starts `app/server.go`.

`server.go` starts new gorilla/mux router, which is filled of routes from `router/router.go`

`router.go` uses routes from `router/routes/routes.go` (and others). Pay your attention of how does middleware used.

Look on `SetupRoutesWithMiddleware method` of `routes.go` file. It uses middleware chain, which is defined in `middleware/middleware_chains.go`. 
You can easily extend this chains for individual routes in Route declaration, just using your own chain or extending default one.

Your controllers may be anywhere, but I generally prefer to place them in `api_controllers` or `controllers` directories. 
