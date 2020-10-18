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

func TestSystemErrorHiddenForProduction(t *testing.T) {
	app := setUp()

	// Given
	app.Bind("config.App.Debug", false)
	var response inter.Response = outcome.NewResponse(outcome.Options{
		App:     app,
		Content: errors.New("incorrect database credentials"),
	})
	decorators := []inter.ResponseDecorator{response_decorator.FilterSensitiveData{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response = bootstrapDecorator.Decorate(response)

	// Then
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"A error has occurred.\"}]}", response.Content())
}

func TestErrorIsLogged(t *testing.T) {
	// Given
	// When
	// Then
}

func TestErrorCode(t *testing.T) {
	// Given
	// When
	// Then
}

func setUp() *foundation.Application {
	return foundation.NewApp()
}