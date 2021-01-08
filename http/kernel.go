package http

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http/http_helper"
	"github.com/confetti-framework/foundation/http/middleware"
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

func (k Kernel) RecoverFromMiddlewarePanic(recover interface{}) inter.Response {
	app := *k.App
	return middleware.NewPipeline(*k.App).
		Send(app.Make("request").(inter.Request)).
		Through(allMiddlewares([]inter.HttpMiddleware{})).
		Then(func(request inter.Request) inter.Response {
			panic(http_helper.GetErrorFromPanic(recover))
		})
}

func allMiddlewares(customMiddlewares []inter.HttpMiddleware) []inter.HttpMiddleware {
	// Append framework middlewares should be placed at the end.
	return append(
		customMiddlewares,
		middleware.DecorateResponse{},
		middleware.PanicToResponse{},
	)
}
