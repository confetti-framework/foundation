package log

import (
	standardErrors "errors"
	"github.com/confetti-framework/contract/inter"
	supportErrors "github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/decorator/response_decorator"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/confetti-framework/foundation/loggers"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/confetti-framework/syslog/log_level"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_error_is_logged_with_standard_error(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger(true)
	responseBefore := newTestResponse(app, standardErrors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response := bootstrapDecorator.Decorate(responseBefore)

	// Then
	require.Same(t, responseBefore, response)
	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Regexp(t, ` \[level severity="emerg"\] incorrect database credentials $`, lines[0][0])
}

func Test_error_without_trace(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger(true)
	responseBefore := newTestResponse(app, supportErrors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	bootstrapDecorator.Decorate(responseBefore)

	// Then
	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Regexp(t, ` \[level severity="emerg"\] incorrect database credentials $`, lines[0][0])
}

func Test_error_trace(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger(false)
	responseBefore := newTestResponse(app, supportErrors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	bootstrapDecorator.Decorate(responseBefore)

	// Then
	lines := openAndReadFile(testFile)
	require.Greater(t, len(lines), 3)
	require.Regexp(t, ` \[level severity="emerg"\] incorrect database credentials $`, lines[0][0])
	if len(lines) > 1 {
		require.Regexp(t, `log.Test_error_trace`, lines[1][0])
		require.Regexp(t, `log/log_error_from_response_test.go:[0-9]+$`, lines[2][0])
	}
}

func Test_don_not_log_ignored_logs(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger(false)
	app.Bind("config.Errors.NoLogging", []interface{}{validationError})
	responseBefore := newTestResponse(app, validationError)
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	bootstrapDecorator.Decorate(responseBefore)

	// Then
	lines := openAndReadFile(testFile)
	require.Len(t, lines, 0)
}

func Test_log_debug_level_from_error(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger(true)
	responseBefore := newTestResponse(app, supportErrors.WithLevel(supportErrors.New("user not found"), log_level.DEBUG))
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	bootstrapDecorator.Decorate(responseBefore)

	// Then
	lines := openAndReadFile(testFile)
	require.Regexp(t, `level severity="debug"`, lines[0][0])
}

func Test_wrap_error(t *testing.T) {
	// Given
	app := setUpAppWithDefaultLogger(true)
	responseBefore := newTestResponse(app, supportErrors.Wrap(standardErrors.New("user id not found"), "validation error"))
	decorators := []inter.ResponseDecorator{response_decorator.LogError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	bootstrapDecorator.Decorate(responseBefore)

	// Then
	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], `validation error: user id not found`)
}

func newTestResponse(app inter.App, content error) inter.Response {
	app.Bind("outcome_json_encoders", mock.JsonEncoders)

	var response inter.Response = outcome.NewResponse(outcome.Options{
		App:      app,
		Content:  content,
		Encoders: "outcome_json_encoders",
	})
	return response
}

func setUpAppWithDefaultLogger(stackTrace bool) inter.App {
	setUp()

	single := loggers.Syslog{Path: testFile, MinLevel: log_level.DEBUG, HideStackTrace: stackTrace}
	allLoggers := map[string]interface{}{"single": single}

	app := newTestApp()
	app.Bind("config.Errors.NoLogging", []interface{}{})
	app.Bind("config.Logging.Channels", allLoggers)
	app.Bind("config.Logging.Default", "single")

	return app
}

var validationError = supportErrors.New("validation error")
