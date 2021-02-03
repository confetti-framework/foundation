package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/method"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/confetti-framework/foundation/http/routing"
	"github.com/stretchr/testify/require"
	"testing"
)

type MockedMiddleware1 struct{}

func (t MockedMiddleware1) Handle(request inter.Request, next inter.Next) inter.Response {
	request.SetBody(request.Body() + "1 -> ")
	response := next(request)
	return response.Body(response.GetBody() + " <- 1")
}

type MockedMiddleware2 struct{}

func (t MockedMiddleware2) Handle(request inter.Request, next inter.Next) inter.Response {
	request.SetBody(request.Body() + "2 -> ")
	response := next(request)
	return response.Body(response.GetBody() + " <- 2")
}

type MockedMiddleware3 struct{}

func (t MockedMiddleware3) Handle(request inter.Request, next inter.Next) inter.Response {
	request.SetBody(request.Body() + "3 -> ")
	response := next(request)
	return response.Body(response.GetBody() + " <- 3")
}

func Test_middleware_on_group(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/roles",
	})

	routes := routing.Group(
		routing.Group(
			routing.Get("/roles", func(request inter.Request) inter.Response {
				return outcome.Html(request.Body())
			}),
		).Middleware(MockedMiddleware3{}),
	).Middleware(
		MockedMiddleware1{},
		MockedMiddleware2{},
	)

	request.App().Singleton("routes", routes)

	// When
	response := http.Kernel{}.Handle(request)

	// Then
	require.Equal(t, "1 -> 2 -> 3 ->  <- 3 <- 2 <- 1", response.GetBody())
}

func Test_without_middleware_on_group(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/roles",
	})

	routes := routing.Group(
		routing.Get("/roles", func(request inter.Request) inter.Response {
			return outcome.Html(request.Body())
		}).WithoutMiddleware(MockedMiddleware1{}),
	).Middleware(
		MockedMiddleware1{},
		MockedMiddleware2{},
	)

	request.App().Singleton("routes", routes)

	// When
	response := http.Kernel{}.Handle(request)

	// Then
	require.Equal(t, "2 ->  <- 2", response.GetBody())
}
