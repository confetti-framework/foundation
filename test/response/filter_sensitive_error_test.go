package response

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/decorator/response_decorator"
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/http/outcome"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/stretchr/testify/require"
	"html/template"
	net "net/http"
	"testing"
)

func Test_system_error_filter_for_production(t *testing.T) {
	app := setUp()

	// Given
	app.Bind("config.App.Debug", false)
	response := newTestResponse(app, errors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.HttpStatus{}, response_decorator.FilterSensitiveError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response = bootstrapDecorator.Decorate(response)

	// Then
	require.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"An error has occurred"}]}`, response.GetBody())
}

func Test_system_error_show_for_development(t *testing.T) {
	app := setUp()

	// Given
	app.Bind("config.App.Debug", true)
	response := newTestResponse(app, errors.New("incorrect database credentials"))
	decorators := []inter.ResponseDecorator{response_decorator.FilterSensitiveError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response = bootstrapDecorator.Decorate(response)

	// Then
	require.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"Incorrect database credentials"}]}`, response.GetBody())
}

func Test_show_user_error(t *testing.T) {
	app := setUp()

	// Given
	app.Bind("config.App.Debug", true)
	response := newTestResponse(app, errors.WithStatus(errors.New("invalid user id"), net.StatusNotFound))
	decorators := []inter.ResponseDecorator{response_decorator.FilterSensitiveError{}}
	bootstrapDecorator := response_decorator.Handler{Decorators: decorators}

	// When
	response = bootstrapDecorator.Decorate(response)

	// Then
	require.Equal(t, `{"jsonapi":{"version":"1.0"},"errors":[{"title":"Invalid user id"}]}`, response.GetBody())
}

func setUp() *foundation.Application {
	app := foundation.NewApp()
	app.Bind("config.App.Debug", false)
	app.Bind("outcome_json_encoders", mock.JsonEncoders)
	app.Bind("outcome_html_encoders", append(
		mock.HtmlEncoders,
		encoder.ErrorsToHtml{View: mock.NewViewErrorMock},
	))
	app.Bind("template_builder", func(template *template.Template) (*template.Template, error) {
		return template.ParseGlob(mock.TemplateByName("") + "/[^sub]*")
	})

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
