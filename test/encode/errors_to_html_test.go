package encode

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_can_not_convert_string_to_html_error(t *testing.T) {
	result := encoder.ErrorsToHtml{}.IsAble("Foo")
	require.False(t, result)
}

func Test_one_error_can_convert_to_html(t *testing.T) {
	result := encoder.ErrorsToHtml{}.IsAble(errors.New("entity not found"))
	require.True(t, result)
}

func Test_slice_with_one_error_convert_to_html(t *testing.T) {
	result := encoder.ErrorsToHtml{}.IsAble([]error{errors.New("entity not found")})
	require.True(t, result)
}

func Test_not_correct_error_can_not_convert_to_html(t *testing.T) {
	app := setUp()
	encoders := []inter.Encoder{encoder.ErrorsToHtml{}}
	encoder := encoder.ErrorsToHtml{}
	result, err := encoder.EncodeThrough(app, "foo", encoders)

	require.Equal(t, "", result)
	require.EqualError(t, err, "can't convert object to html in error format")
}

func Test_one_error_to_html_without_template(t *testing.T) {
	app := setUp()
	app.Bind("config.App.Debug", false)
	encoder := encoder.ErrorsToHtml{}
	result, err := encoder.EncodeThrough(app, errors.New("entity not found"), mock.HtmlEncoders)

	require.NoError(t, err)
	require.Equal(t, "Entity not found", result)
	require.NotContains(t, result, "<p>")
	require.NotContains(t, result, "errors_to_html.go")
}

func Test_one_error_to_html_on_production(t *testing.T) {
	app := setUp()
	app.Bind("config.App.Debug", false)
	encoder := encoder.ErrorsToHtml{View: mock.NewViewErrorMock}
	result, err := encoder.EncodeThrough(app, errors.New("entity not found"), mock.HtmlEncoders)

	require.NoError(t, err)
	require.Equal(t, "<h1>500</h1>\n<h2>Entity not found</h2>\n", result)
	require.NotContains(t, result, "<p>")
	require.NotContains(t, result, "errors_to_html.go")
}

func Test_one_error_to_html_on_development_with_stack_trace(t *testing.T) {
	app := setUp()
	app.Bind("config.App.Debug", true)
	encoder := encoder.ErrorsToHtml{View: mock.NewViewErrorMock}
	result, err := encoder.EncodeThrough(app, errors.New("entity not found"), mock.HtmlEncoders)

	require.NoError(t, err)
	require.Contains(t, result, "Entity not found")
	require.Contains(t, result, "errors_to_html_test.go")
}

func Test_template_error_to_html(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorsToHtml{View: mock.NewViewErrorMock}.
		EncodeThrough(
			app,
			errors.New("entity not found"),
			mock.HtmlEncoders,
		)

	require.NoError(t, err)
	require.Equal(t, "<h1>500</h1>\n<h2>Entity not found</h2>\n", result)
}

func Test_slice_without_error_to_html(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorsToHtml{View: mock.NewViewErrorMock}.
		EncodeThrough(
			app,
			[]error{},
			mock.HtmlEncoders,
		)

	require.Equal(t, "", result)
	require.EqualError(t, err, "can't convert object to html in error format")
}

func Test_slice_with_one_error_to_html(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorsToHtml{View: mock.NewViewErrorMock}.
		EncodeThrough(
			app,
			[]error{errors.New("entity not found")},
			mock.HtmlEncoders,
		)

	require.NoError(t, err)
	require.Equal(t, "<h1>500</h1>\n<h2>Entity not found</h2>\n", result)
}

func Test_system_error_convert_to_html(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, errors.New("bad request"), []inter.Encoder{encoder.InterfaceToHtml{}})

	require.Equal(t, "no encoder found to handle error: bad request", result)
	require.EqualError(t, err, "no encoder found to handle error: bad request")
}
