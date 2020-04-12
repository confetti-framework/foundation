package http

import (
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http/lanvard"
	"github.com/lanvard/routing/router"
	"lanvard/app/http/decorator"
	"lanvard/app/http/middleware"
)

type Kernel struct {
	App        *foundation.Application
	Router     router.Router
	Middleware []middleware.PipeInterface
}

// Handle an incoming HTTP request.
func (k Kernel) Handle(request lanvard.Request) lanvard.Response {
	return k.sendRequestThroughRouter(request)
	// @todo event RequestHandled
}

// Send the given request through the middleware / router.
func (k Kernel) sendRequestThroughRouter(request lanvard.Request) lanvard.Response {
	k.App.Container().Instance("request", request)

	return middleware.NewPipeline(k.App).
		Send(request).
		Through(k.Middleware).
		Then(k.dispatchToRouter())
}

func (k Kernel) Bootstrap() *foundation.Application {
	k.App = decorator.Bootstrap(k.App)
	k.App.HasBeenBootstrapped = true

	return k.App
}

func (k Kernel) dispatchToRouter() middleware.Destination {
	return func(request lanvard.Request) lanvard.Response {
		return router.NewRouter(k.App).DispatchToRoute(request)
	}
}
