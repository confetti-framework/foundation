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

func Test_method_match_with_get(t *testing.T) {
	// Given
	routes := routing.Match([]string{"GET", "POST"}, "/users", emptyController())

	// When
	route := routes.Match(newRequest(http.Options{
		Method: method.Get,
		Url:    "/users",
	}))

	// Then
	require.NotNil(t, route)
}

func Test_method_match_with_second_method(t *testing.T) {
	// Given
	routes := routing.Match([]string{method.Get, method.Post}, "/users", emptyController())

	// When
	route := routes.Match(newRequest(http.Options{
		Method: method.Post,
		Url:    "/users",
	}))

	// Then
	require.NotNil(t, route)
}

func Test_fallback_controller(t *testing.T) {
	// Given
	routes := routing.Group(
		routing.Get("/test", emptyController()),
		routing.Fallback(func(request inter.Request) inter.Response {
			return outcome.Html("404 Page not found")
		}),
	)

	// When
	request := newRequest(http.Options{Method: method.Get, Url: ""})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "404 Page not found", response.GetBody())
}
