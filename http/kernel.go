package http

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/middleware"
)

type Kernel struct {
	App        *inter.App
	Router     Router
	Middleware []inter.Pipe
}

// Handle an incoming HTTP request.
func (k Kernel) Handle(request inter.Request) inter.Response {
	return k.sendRequestThroughRouter(request)
	// @todo event RequestHandled
}

// Send the given request through the middleware / router.
func (k Kernel) sendRequestThroughRouter(request inter.Request) inter.Response {
	(*k.App).Instance("request", request)


	// todo bind in new app
	return middleware.NewPipeline(k.App).
		Send(request).
		Through(k.Middleware).
		Then(k.dispatchToRouter())
}

func (k Kernel) dispatchToRouter() inter.MiddlewareDestination {
	return func(request inter.Request) inter.Response {
		return NewRouter(k.App).DispatchToRoute(request)
	}
}
