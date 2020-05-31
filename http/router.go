package http

import (
	"github.com/lanvard/contract/inter"
)

type Router struct {
	app            *inter.App
	routes         inter.RouteCollection
	currentRequest inter.Request
}

func NewRouter(appPointer *inter.App) Router {
	app := *appPointer
	routes := app.Make("routes").(inter.RouteCollection)

	return Router{app: &app, routes: routes}
}

func (r Router) DispatchToRoute(request inter.Request) inter.Response {
	r.currentRequest = request

	route := r.routes.Match(request)

	// todo implement event Events\RouteMatched
	// todo implement RouteMiddleware
	return route.Controller()(request)
}
