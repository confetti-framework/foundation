package log

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/decorator/response_decorator"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorIsLogged(t *testing.T) {
	setUp()

	// Given
	responseBefore := newTestResponse(errors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response := bootstrapDecorator.Decorate(responseBefore)

	// Then
	assert.Same(t, responseBefore, response)

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Regexp(t, ` \[level severity="crit"\] incorrect database credentials {"key":"value"}$`, lines[0][0])

}

func TestErrorCode(t *testing.T) {
	// Given
	// When
	// Then
}

func TestErrorTrace(t *testing.T) {
	// Given
	// When
	// Then
}

func newTestResponse(content error) inter.Response {
	app := foundation.NewApp()
	app.Bind("outcome_json_encoders", outcome.JsonEncoders)

	var response inter.Response = outcome.NewResponse(outcome.Options{
		App:      app,
		Content:  content,
		Encoders: "outcome_json_encoders",
	})
	return response
}
