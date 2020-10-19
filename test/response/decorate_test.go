package response

import (
	"github.com/lanvard/foundation"
	"testing"
)

func TestSystemErrorHiddenForProduction(t *testing.T) {
	// app := setUp()
	//
	// // Given
	// app.Bind("config.App.Debug", false)
	// var response inter.Response = outcome.NewResponse(outcome.Options{
	// 	App:     app,
	// 	Content: errors.New("incorrect database credentials"),
	// })
	// decorators := []inter.ResponseDecorator{response_decorator.FilterSensitiveData{}}
	// bootstrapDecorator := response_decorator.Handler{Decorators: decorators}
	//
	// // When
	// response = bootstrapDecorator.Decorate(response)
	//
	// // Then
	// assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"A error has occurred.\"}]}", response.Body())
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
