package log

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/decorator/response_decorator"
	"github.com/lanvard/foundation/loggers"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/syslog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorIsLogged(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger()
	responseBefore := newTestResponse(app, errors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response := bootstrapDecorator.Decorate(responseBefore)

	// Then
	assert.Same(t, responseBefore, response)

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Regexp(t, ` \[level severity="emerg"\] incorrect database credentials $`, lines[0][0])
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

func newTestResponse(app inter.App, content error) inter.Response {
	app.Bind("outcome_json_encoders", outcome.JsonEncoders)

	var response inter.Response = outcome.NewResponse(outcome.Options{
		App:      app,
		Content:  content,
		Encoders: "outcome_json_encoders",
	})
	return response
}

func setUpAppWithDefaultLogger() inter.App {
	setUp()

	single := loggers.Syslog{Path: testFile, MinLevel: syslog.DEBUG}
	allLoggers := map[string]interface{}{"single": single}

	app := newTestApp()
	app.Bind("config.Logging.Channels", allLoggers)
	app.Bind("config.Logging.Default", "single")

	return app
}
