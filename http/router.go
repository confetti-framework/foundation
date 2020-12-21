package http

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/middleware"
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

	middlewares := route.Middleware()

	// Framework middlewares should be placed at the end
	// so that they are executed first when a response is returned
	middlewares = append(
		middlewares,
		middleware.DecorateResponse{},
		middleware.AppendAppMiddleware{},
		middleware.PanicToResponse{},
	)

	return middleware.NewPipeline(request.App()).
		Send(request).
		Through(middlewares).
		Then(route.Controller())
}
