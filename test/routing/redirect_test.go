package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/method"
	"github.com/confetti-framework/foundation/http/outcome"
	routing2 "github.com/confetti-framework/foundation/http/routing"
	"github.com/stretchr/testify/require"
	net "net/http"
	"testing"
)

func Test_simple_redirect(t *testing.T) {
	response := outcome.Redirect("home/dashboard", net.StatusPermanentRedirect)
	require.Equal(t, "home/dashboard", response.GetHeader("Location"))
}

func Test_redirect_route(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/here",
	})

	routes := routing2.Redirect("/here", "/there", net.StatusFound)

	// When
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, net.StatusFound, response.GetStatus())
	require.Contains(t, response.GetBody(), "Redirecting to <a href=\"/there\">")
	require.Equal(t, "/there", response.GetHeader("Location"))
}

func Test_redirect_temporary_route(t *testing.T) {
	// Given
	routes := routing2.RedirectTemporary("/here", "/there")

	// When
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/here",
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, net.StatusFound, response.GetStatus())
}

func Test_redirect_permanent_route(t *testing.T) {
	// Given
	routes := routing2.RedirectPermanent("/here", "/there")

	// When
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/here",
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, net.StatusMovedPermanently, response.GetStatus())
}

func Test_redirect_by_named_route(t *testing.T) {
	// Given
	routes := routing2.Group(
		routing2.Get("/roles", emptyController()).Name("roles"),
		routing2.Get("/comments", func(request inter.Request) inter.Response {
			return outcome.RedirectToRoute(request.App(), "roles")
		}).Name("Comments"),
	)

	// When
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/comments",
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, net.StatusFound, response.GetStatus())
	require.Contains(t, response.GetBody(), "/roles")
}

func Test_redirect_by_named_route_with_parameters(t *testing.T) {
	// Given
	routes := routing2.Group(
		routing2.Get("/roles/{id}", emptyController()).Name("roles"),
		routing2.Get("/comments", func(request inter.Request) inter.Response {
			return outcome.RedirectToRoute(
				request.App(),
				"roles",
				outcome.Parameters{"id": 12},
			)
		}).Name("Comments"),
	)

	// When
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/comments",
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, net.StatusFound, response.GetStatus())
	require.Contains(t, response.GetBody(), "/roles/12")
}

func Test_redirect_custom(t *testing.T) {
	routeCollection := routing2.Redirect("/foo", "/bar", 200)

	route := routeCollection.All()[0]

	require.Equal(t, "/foo", route.Uri())
	require.Equal(t, "GET", route.Method())
	require.Equal(t, 200, route.RouteOptions().Status())
}

func Test_redirect_permanent(t *testing.T) {
	routeCollection := routing2.RedirectPermanent("/foo", "/bar")

	route := routeCollection.All()[0]

	require.Equal(t, "/foo", route.Uri())
	require.Equal(t, "GET", route.Method())
	require.Equal(t, 301, route.RouteOptions().Status())
}

func Test_redirect_temporary(t *testing.T) {
	routeCollection := routing2.RedirectTemporary("/foo", "/bar")

	route := routeCollection.All()[0]

	require.Equal(t, "/foo", route.Uri())
	require.Equal(t, "GET", route.Method())
	require.Equal(t, 302, route.RouteOptions().Status())
}
