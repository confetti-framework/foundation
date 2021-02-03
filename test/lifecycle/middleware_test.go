package lifecycle

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/middleware"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_middleware_with_app_from_response(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	middlewares := []inter.HttpMiddleware{checkAppRequiredInMiddleware{}, emptyMiddleware{}}

	response := middleware.NewPipeline(request.App()).
		Send(request).
		Through(middlewares).
		Then(func(request inter.Request) inter.Response {
			return outcome.Html("foo")
		})

	require.NotNil(t, response.App())
}

func Test_middleware_with_app_with_response_in_middleware(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	middlewares := []inter.HttpMiddleware{checkAppRequiredInMiddleware{}, middlewareWithResponse{}, emptyMiddleware{}}

	response := middleware.NewPipeline(request.App()).
		Send(request).
		Through(middlewares).
		Then(func(request inter.Request) inter.Response {
			return outcome.Html("foo")
		})

	require.NotNil(t, response.App())
}

type checkAppRequiredInMiddleware struct{}

func (c checkAppRequiredInMiddleware) Handle(request inter.Request, next inter.Next) inter.Response {
	response := next(request)
	if response.App() == nil {
		panic("app is null")
	}

	return response
}

type emptyMiddleware struct{}

func (e emptyMiddleware) Handle(request inter.Request, next inter.Next) inter.Response {
	return next(request)
}

type middlewareWithResponse struct{}

func (m middlewareWithResponse) Handle(request inter.Request, next inter.Next) inter.Response {
	return outcome.Html("test")
}
