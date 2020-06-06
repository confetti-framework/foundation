package http

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/middleware"
)

type Router struct {
	routes         inter.RouteCollection
	currentRequest inter.Request
}

func NewRouter(app inter.App) Router {
	routes := app.Make("routes").(inter.RouteCollection)

	return Router{routes: routes}
}

func (r Router) DispatchToRoute(request inter.Request) inter.Response {
	r.currentRequest = request

	route := r.routes.Match(request)

	// todo implement event Events\RouteMatched

	return middleware.NewPipeline(request.App()).
		Send(request).
		Through(route.Middleware()).
		Then(route.Controller())
}
