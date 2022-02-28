package response

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/decorator/response_decorator"
	"github.com/confetti-framework/foundation/http"
	"github.com/confetti-framework/foundation/http/middleware"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/stretchr/testify/require"
	net "net/http"
	"testing"
)

var userNotFound = errors.New("user not found")

func Test_panic_without_valid_encoder_defined(t *testing.T) {
	// Given
	app := setUp()
	request := http.NewRequest(http.Options{App: app})

	// When
	when := func() {
		middleware.PanicToResponse{}.Handle(
			request,
			func(request inter.Request) inter.Response {
				panic(userNotFound)
			},
		)
	}

	// Then
	require.Panics(t, when)
}

func Test_panic_error_to_json(t *testing.T) {
	// Given
	app := setUp()
	app.Bind("default_response_outcome", outcome.Json)
	request := http.NewRequest(http.Options{App: app})

	// When
	response := middleware.PanicToResponse{}.Handle(
		request,
		func(request inter.Request) inter.Response {
			panic(userNotFound)
		},
	)
	response.SetApp(request.App())

	// Then
	require.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"User not found"}]}`, response.GetBody())
}

func Test_panic_string_to_json(t *testing.T) {
	// Given
	app := setUp()
	app.Bind("default_response_outcome", outcome.Json)
	request := http.NewRequest(http.Options{App: app})

	// When
	response := middleware.PanicToResponse{}.Handle(
		request,
		func(request inter.Request) inter.Response {
			panic(userNotFound)
		},
	)
	response.SetApp(request.App())

	// Then
	require.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"User not found"}]}`, response.GetBody())
}

func Test_panic_unknown_to_json(t *testing.T) {
	// Given
	app := setUp()
	app.Bind("default_response_outcome", outcome.Json)
	request := http.NewRequest(http.Options{App: app})

	// When
	response := middleware.PanicToResponse{}.Handle(
		request,
		func(request inter.Request) inter.Response {
			panic(invalidError{})
		},
	)
	response.SetApp(request.App())

	// Then
	require.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"Can't convert panic to response. Error or string required"}]}`, response.GetBody())
}

func Test_panic_with_error_to_html(t *testing.T) {
	// Given
	app := setUp()
	app.Bind("default_response_outcome", outcome.Html)
	request := http.NewRequest(http.Options{App: app})

	// When
	response := middleware.PanicToResponse{}.Handle(
		request,
		func(request inter.Request) inter.Response {
			panic(userNotFound)
		},
	)
	response.SetApp(request.App())

	// Then
	require.Equal(t, "<h1>500</h1>\n<h2>User not found</h2>\n", response.GetBody())
}

func Test_http_status_to_response(t *testing.T) {
	// Given
	app := setUp()
	responseBefore := newTestResponse(app, errors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.HttpStatus{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response := bootstrapDecorator.Decorate(responseBefore)

	// Then
	require.Equal(t, net.StatusInternalServerError, response.GetStatus())
}

func Test_http_status_bad_request_to_response(t *testing.T) {
	// Given
	app := setUp()
	responseBefore := newTestResponse(app, errors.WithStatus(errors.New("incorrect database credentials"), net.StatusBadRequest))
	decorators := []inter.ResponseDecorator{response_decorator.HttpStatus{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response := bootstrapDecorator.Decorate(responseBefore)

	// Then
	require.Equal(t, net.StatusBadRequest, response.GetStatus())
}

type invalidError struct{}
