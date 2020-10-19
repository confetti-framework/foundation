package response

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/middleware"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorToJson(t *testing.T) {
	// Given
	app := setUp()
	request := http.NewRequest(http.Options{App: app})

	// When
	response := middleware.PanicToResponse{Outcome: outcome.Json}.Handle(
		request,
		func(request inter.Request) inter.Response {
			panic(simpleError{})
		},
	)
	response.SetApp(request.App())

	// Then
	assert.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"No user found"}]}`, response.Body())
}

func TestStringToJson(t *testing.T) {
	// Given
	app := setUp()
	request := http.NewRequest(http.Options{App: app})

	// When
	response := middleware.PanicToResponse{Outcome: outcome.Json}.Handle(
		request,
		func(request inter.Request) inter.Response {
			panic("no user found")
		},
	)
	response.SetApp(request.App())

	// Then
	assert.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"No user found"}]}`, response.Body())
}

func TestUnknownToJson(t *testing.T) {
	// Given
	app := setUp()
	request := http.NewRequest(http.Options{App: app})

	// When
	response := middleware.PanicToResponse{Outcome: outcome.Json}.Handle(
		request,
		func(request inter.Request) inter.Response {
			panic(invalidError{})
		},
	)
	response.SetApp(request.App())

	// Then
	assert.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"Can't convert panic to response. Error or string required"}]}`, response.Body())
}

type simpleError struct{}

func (s simpleError) Error() string {
	return "no user found"
}

type invalidError struct{}
