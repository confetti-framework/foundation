package response

import (
	standardErrors "errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/decorator/response_decorator"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/middleware"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	net "net/http"
	"testing"
)

func TestPanicWithoutValidEncoderDefined(t *testing.T) {
	// Given
	app := setUp()
	request := http.NewRequest(http.Options{App: app})

	// When
	when := func() {
		middleware.PanicToResponse{}.Handle(
			request,
			func(request inter.Request) inter.Response {
				panic(simpleError{})
			},
		)
	}

	// Then
	assert.Panics(t, when)
}

func TestPanicErrorToJson(t *testing.T) {
	// Given
	app := setUp()
	app.Bind("response_encoder", outcome.Json)
	request := http.NewRequest(http.Options{App: app})

	// When
	response := middleware.PanicToResponse{}.Handle(
		request,
		func(request inter.Request) inter.Response {
			panic(simpleError{})
		},
	)
	response.SetApp(request.App())

	// Then
	assert.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"No user found"}]}`, response.Body())
}

func TestPanicStringToJson(t *testing.T) {
	// Given
	app := setUp()
	app.Bind("response_encoder", outcome.Json)
	request := http.NewRequest(http.Options{App: app})

	// When
	response := middleware.PanicToResponse{}.Handle(
		request,
		func(request inter.Request) inter.Response {
			panic("no user found")
		},
	)
	response.SetApp(request.App())

	// Then
	assert.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"No user found"}]}`, response.Body())
}

func TestPanicUnknownToJson(t *testing.T) {
	// Given
	app := setUp()
	app.Bind("response_encoder", outcome.Json)
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
	assert.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"Can't convert panic to response. Error or string required"}]}`, response.Body())
}

func TestHttpStatusToResponse(t *testing.T) {
	// Given
	app := setUp()
	responseBefore := newTestResponse(app, standardErrors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.HttpStatus{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response := bootstrapDecorator.Decorate(responseBefore)

	// Then
	assert.Equal(t, net.StatusInternalServerError, response.Status())
}

type simpleError struct{}

func (s simpleError) Error() string {
	return "no user found"
}

type invalidError struct{}
