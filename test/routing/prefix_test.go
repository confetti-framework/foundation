package routing

import (
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/method"
	"github.com/confetti-framework/foundation/http/routing"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_prefix_url(t *testing.T) {
	// Given
	routes := routing.Get("/user/{user_id}/roles", emptyController()).Prefix("/api")

	// When
	matchedRoute := routes.Match(newRequest(http.Options{
		Method: method.Get,
		Url:    "/api/user/{user_id}/roles",
	}))

	// Then
	require.NotNil(t, matchedRoute)
}

func Test_multiple_prefixes(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/admin/user/{user_id}/roles",
	})

	routeGroup := routing.Group(
		routing.Get("/{user_id}/roles", emptyController()).Prefix("/user"),
	).Prefix("/admin")

	// When
	matchedRoute := routeGroup.Match(request)

	// Then
	require.NotNil(t, matchedRoute)
}
