package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/method"
	"github.com/confetti-framework/foundation/http/outcome"
	routing2 "github.com/confetti-framework/foundation/http/routing"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_regular_expression_after_name(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/user/Bla",
	})

	// When
	routes := routing2.Get("/user/{username:[A-Za-z]+}", func(request inter.Request) inter.Response {
		return outcome.Html("value: " + request.Parameter("username").String())
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "value: Bla", response.GetBody())
}

func Test_global_constraint(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/user/bla",
	})

	// When
	routes := routing2.Get("/user/{username}", func(request inter.Request) inter.Response {
		return outcome.Html("value: " + request.Parameter("username").String())
	}).Where("username", "[a-z]+")
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "value: bla", response.GetBody())
}

func Test_global_constraints(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/user/piet",
	})

	// When
	routes := routing2.Get("/user/{username}", func(request inter.Request) inter.Response {
		return outcome.Html("value: " + request.Parameter("username").String())
	}).WhereMulti(map[string]string{
		"username": "[a-z]+",
	})

	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "value: piet", response.GetBody())
}

func Test_not_valid_global_constraints(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/user/123",
	})

	routes := routing2.Get("/user/{username}", emptyController())

	// When
	when := func() inter.Route {
		return routes.
			WhereMulti(map[string]string{
				"username": "[a-z]+",
			}).
			Match(request)
	}

	// Then
	response := when().Controller()(request)
	require.Equal(t, routing2.RouteNotFoundError, response.GetContent())
}
