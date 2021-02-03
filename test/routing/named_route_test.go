package routing

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/method"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/confetti-framework/foundation/http/routing"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_named_route(t *testing.T) {
	routeCollection := routing.Get("/users/roles", emptyController()).Name("Roles")
	route := routeCollection.All()[0]

	require.True(t, route.Named("Roles"))
}

func Test_multiple_named_routes(t *testing.T) {
	// Given
	routes := routing.Group(
		routing.Get("/roles", emptyController()).Name("Roles"),
		routing.Get("/comments", emptyController()).Name("Comments"),
	).Prefix("/users").Name("Users")

	// When
	route := routes.Match(newRequest(http.Options{
		Method: method.Get,
		Url:    "/users/roles",
	}))

	// Then
	require.True(t, route.Named("UsersRoles"))
}

func Test_no_named_route(t *testing.T) {
	routeCollection := routing.Get("/users/roles", emptyController()).Name("Users")
	route := routeCollection.All()[0]

	require.False(t, route.Named("Roles"))
}

func Test_full_url_by_name(t *testing.T) {
	// Given
	routes := routing.Get("/user/{user_id}/type/{type}", emptyController()).Domain("{account_name}.endless.horse").Name("Users")

	app := foundation.NewApp()
	app.Singleton("routes", routes)

	// When
	url := outcome.UrlByName(
		app,
		"Users",
		outcome.Parameters{
			"user_id":      12,
			"type":         "author",
			"account_name": "big",
		},
		outcome.Parameters{
			"order_by": "name",
			"size":     50,
		},
	)

	// Then
	// Golang's map has no specific order. The order of the parameters cannot be tested.
	require.Contains(t, url, "https://big.endless.horse/user/12/type/author?")
	require.Contains(t, url, "order_by=name")
	require.Contains(t, url, "size=50")
}

func Test_url_by_name_not_found(t *testing.T) {
	// Given
	routes := routing.Group(
		routing.Get("/users", emptyController()).Name("Users"),
		routing.Get("/comments", emptyController()).Name("Comments"),
	)

	app := foundation.NewApp()
	app.Singleton("routes", routes)

	// When
	when := func() {
		outcome.UrlByName(app, "Roles")
	}

	// Then
	require.Panics(t, when, "URL cannot be generated because no route was found matching name Roles")
}

func Test_url_by_name_multiple_matches_found(t *testing.T) {
	// Given
	routes := routing.Group(
		routing.Get("/roles", emptyController()).Name("UsersRole"),
		routing.Get("/comments", emptyController()).Name("UsersComments"),
	)

	app := foundation.NewApp()
	app.Singleton("routes", routes)

	// When
	when := func() {
		outcome.UrlByName(app, "Users")
	}

	// Then
	require.Panics(t, when, "URL cannot be generated because multiple routes were found matching name Users")
}

func Test_url_by_name_with_optional_parameter(t *testing.T) {
	// Given
	routes := routing.Group(
		routing.Get("/roles/{role?}", emptyController()).Name("UsersRole"),
		routing.Get("/comments", emptyController()).Name("UsersComments"),
	)

	app := foundation.NewApp()
	app.Singleton("routes", routes)

	// When
	when := func() {
		outcome.UrlByName(app, "Users")
	}

	// Then
	require.Panics(t, when, "URL cannot be generated because multiple routes were found matching name Users")
}

func Test_route_match_from_request(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/roles",
	})

	routes := routing.Group(
		routing.Get("/roles", func(request inter.Request) inter.Response {
			if request.Route().Named("Roles") {
				return outcome.Html("Has match")
			} else {
				return outcome.Html("No match")
			}
		}).Name("Roles"),
		routing.Get("/comments", emptyController()).Name("Comments"),
	)

	// When
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "Has match", response.GetBody())
}

func Test_route_no_match_from_request(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/roles",
	})

	routes := routing.Group(
		routing.Get("/roles", func(request inter.Request) inter.Response {
			if request.Route().Named("Comments") {
				return outcome.Html("Has match")
			} else {
				return outcome.Html("No match")
			}
		}).Name("Roles"),
		routing.Get("/comments", emptyController()).Name("Comments"),
	)

	// When
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "No match", response.GetBody())
}
