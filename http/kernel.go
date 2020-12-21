package http

import (
	"github.com/confetti-framework/contract/inter"
)

type Kernel struct {
	App *inter.App
}

// Handle an incoming HTTP request.
func (k Kernel) Handle(request inter.Request) inter.Response {
	return k.sendRequestThroughRouter(request)
}

// Send the given request through the middleware / router.
func (k Kernel) sendRequestThroughRouter(request inter.Request) inter.Response {
	request.App().Bind("request", request)

	return NewRouter(request.App()).DispatchToRoute(request)
}
