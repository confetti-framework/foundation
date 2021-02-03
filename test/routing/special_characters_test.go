package routing

import (
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/method"
	routing2 "github.com/confetti-framework/foundation/http/routing"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_without_begin_slash(t *testing.T) {
	// Given
	routes := routing2.Get("user/roles/", emptyController())

	// When
	matchedRoute := routes.Match(newRequest(http.Options{
		Method: method.Get,
		Url:    "/user/roles/",
	}))

	// Then
	require.Equal(t, "/user/roles/", matchedRoute.Uri())
}

func Test_without_end_slash(t *testing.T) {
	// Given
	routes := routing2.Get("/user/roles", emptyController())

	// When
	matchedRoute := routes.Match(newRequest(http.Options{
		Method: method.Get,
		Url:    "/user/roles",
	}))

	// Then
	require.Equal(t, "/user/roles{allow_slash:\\/?}", matchedRoute.Uri())
}

func Test_with_slash(t *testing.T) {
	// Given
	routes := routing2.Get("/users", emptyController())

	// When
	matchedRoute := routes.Match(newRequest(http.Options{
		Method: method.Get,
		Url:    "/users/",
	}))

	// Then
	require.Equal(t, "/users{allow_slash:\\/?}", matchedRoute.Uri())
}

func Test_with_slash_and_query_parameters(t *testing.T) {
	// Given
	routes := routing2.Get("/users", emptyController())

	// When
	matchedRoute := routes.Match(newRequest(http.Options{
		Method: method.Get,
		Url:    "/users/?bla",
	}))

	// Then
	require.Equal(t, "/users{allow_slash:\\/?}", matchedRoute.Uri())
}
