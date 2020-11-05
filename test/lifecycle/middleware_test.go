package lifecycle

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/middleware"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddlewareWithAppFromRequest(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	middlewares := []inter.HttpMiddleware{checkAppRequiredInMiddleware{}, middleware.AppendAppMiddleware{}}

	response := middleware.NewPipeline(request.App()).
		Send(request).
		Through(middlewares).
		Then(func(request inter.Request) inter.Response {
			return outcome.Html("foo")
		})

	assert.NotNil(t, response)
}

type checkAppRequiredInMiddleware struct{}

func (a checkAppRequiredInMiddleware) Handle(request inter.Request, next inter.Next) inter.Response {
	response := next(request)
	if response.App() == nil {
		panic("app is not present")
	}

	return response
}
