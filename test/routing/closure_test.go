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

func Test_name_from_nil(t *testing.T) {
	// Given
	request := newRequest(http.Options{
		Method:  method.Post,
		Url:     "/users",
		Content: "body from request",
	})

	// When
	routes := routing2.Post("/users", func(request inter.Request) inter.Response {
		return outcome.Html(request.Body() + " & body from response")
	})
	result := routes.Match(request).Controller()(request)
	result.SetApp(request.App())

	// Then
	require.Equal(t, "body from request & body from response", result.GetBody())
}
