package response

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/decorator/response_decorator"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSystemErrorFilterForProduction(t *testing.T) {
	app := setUp()

	// Given
	app.Bind("config.App.Debug", false)
	response := newTestResponse(app, errors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.FilterSensitiveError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response = bootstrapDecorator.Decorate(response)

	// Then
	assert.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"An error has occurred"}]}`, response.Body())
}

func TestSystemErrorShowForDevelopment(t *testing.T) {
	app := setUp()

	// Given
	app.Bind("config.App.Debug", true)
	response := newTestResponse(app, errors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.FilterSensitiveError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response = bootstrapDecorator.Decorate(response)

	// Then
	assert.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"Incorrect database credentials"}]}`, response.Body())
}

func setUp() *foundation.Application {
	app := foundation.NewApp()
	app.Bind("outcome_json_encoders", outcome.JsonEncoders)
	return app
}

func newTestResponse(app *foundation.Application, content error) inter.Response {
	var response inter.Response = outcome.NewResponse(outcome.Options{
		App:      app,
		Content:  content,
		Encoders: "outcome_json_encoders",
	})
	return response
}
