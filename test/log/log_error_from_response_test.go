package log

import (
	standardErrors "errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/decorator/response_decorator"
	"github.com/lanvard/foundation/loggers"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/support"
	"github.com/lanvard/syslog/level"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorIsLogged(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger(false)
	responseBefore := newTestResponse(app, standardErrors.New("incorrect database credentials"))
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

func TestErrorTrace(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger(true)
	responseBefore := newTestResponse(app, errors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	bootstrapDecorator.Decorate(responseBefore)

	// Then
	lines := openAndReadFile(testFile)
	assert.Greater(t, len(lines), 3)
	assert.Regexp(t, ` \[level severity="emerg"\] incorrect database credentials $`, lines[0][0])
	assert.Regexp(t, `log.TestErrorTrace$`, lines[1][0])
	assert.Regexp(t, `log_error_from_response_test.go:[0-9]+$`, lines[2][0])
}

func TestLogDebugLevelFromError(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger(true)
	responseBefore := newTestResponse(app, support.NewError("User not found").Level(level.DEBUG))
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	bootstrapDecorator.Decorate(responseBefore)

	// Then
	lines := openAndReadFile(testFile)
	assert.Regexp(t, `level severity="debug"`, lines[0][0])
}

func TestErrorHttpStatus(t *testing.T) {
	// Given
	// When
	// Then
}

func TestErrorCode(t *testing.T) {
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

func setUpAppWithDefaultLogger(stackTrace bool) inter.App {
	setUp()

	single := loggers.Syslog{Path: testFile, MinLevel: level.DEBUG}
	allLoggers := map[string]interface{}{"single": single}

	app := newTestApp()
	app.Bind("config.Logging.Channels", allLoggers)
	app.Bind("config.Logging.Default", "single")

	return app
}
