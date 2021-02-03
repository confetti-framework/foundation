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

func Test_optional_parameter_with_slash(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/user/",
	})

	// When
	routes := routing2.Get("/user/{username?}", func(request inter.Request) inter.Response {
		return outcome.Html("value: " + request.Parameter("username").String())
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "value: ", response.GetBody())
}

func Test_not_match_optional_slash_with_latest_character(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/users",
	})
	routes := routing2.Get("/user/{username?}", emptyController())

	// When
	when := func() inter.Route {
		return routes.Match(request)
	}

	// Then
	response := when().Controller()(request)
	require.Equal(t, routing2.RouteNotFoundError, response.GetContent())
}

func Test_optional_parameter_with_word(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/user-",
	})

	// When
	routes := routing2.Get("/user-{username?}", func(request inter.Request) inter.Response {
		return outcome.Html("value: " + request.Parameter("username").String())
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "value: ", response.GetBody())
}

func Test_multiple_optional_parameter(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/user/jacob/comment/",
	})

	// When
	routes := routing2.Get("/user/{username?}/comment/{command_id?}", func(request inter.Request) inter.Response {
		return outcome.Html("value: " + request.Parameter("username").String())
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "value: jacob", response.GetBody())
}
