package lifecycle

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/middleware"
	"github.com/confetti-framework/routing/outcome"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_middleware_with_app_from_request(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	middlewares := []inter.HttpMiddleware{checkAppRequiredInMiddleware{}, middleware.AppendApp{}}

	response := middleware.NewPipeline(request.App()).
		Send(request).
		Through(middlewares).
		Then(func(request inter.Request) inter.Response {
			return outcome.Html("foo")
		})

	require.NotNil(t, response)
}

type checkAppRequiredInMiddleware struct{}

func (a checkAppRequiredInMiddleware) Handle(request inter.Request, next inter.Next) inter.Response {
	response := next(request)
	if response.App() == nil {
		panic("app is not present")
	}

	return response
}
