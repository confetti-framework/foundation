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

func Test_match_domain(t *testing.T) {
	// Given
	routes := routing.Group(
		routing.Get("/users", emptyController()).Domain("api.bassie.com").Name("withDomain"),
		routing.Get("/users", emptyController()).Name("withoutDomain"),
	)

	// When
	route := routes.Match(newRequest(http.Options{
		Method: method.Get,
		Url:    "/users",
		Host:   "api.bassie.com",
	}))

	// Then
	require.Truef(t, route.Named("withDomain"), "matched with "+route.Name())
}

func Test_don_not_match_domain(t *testing.T) {
	// Given
	routes := routing.Group(
		routing.Get("/users", emptyController()).Domain("api.enadriaan.com").Name("withDomain"),
		routing.Get("/users", emptyController()).Name("withoutDomain"),
	)

	// When
	route := routes.Match(newRequest(http.Options{
		Method: method.Get,
		Url:    "/users",
		Host:   "api.application.com",
	}))

	// Then
	require.Truef(t, route.Named("withoutDomain"), "matched with "+route.Name())
}

func Test_Parameter_from_domain(t *testing.T) {
	// Given
	routes := routing.Group(
		routing.Get("/users", func(request inter.Request) inter.Response {
			return outcome.Html(request.Parameter("account"))
		}),
	).Domain("{account}.bassie.com")

	// When
	request := newRequest(http.Options{
		Method: method.Get,
		Url:    "/users",
		Host:   "klaas.bassie.com",
	})
	response := routes.Match(request).Controller()(request)
	response.SetApp(request.App())

	// Then
	body := response.GetBody()
	require.Equal(t, "klaas", body)
}
