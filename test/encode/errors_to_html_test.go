package encode

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanNotConvertStringToHtmlError(t *testing.T) {
	result := encoder.ErrorToHtml{}.IsAble("Foo")
	assert.False(t, result)
}

func TestOneErrorCanConvertToHtml(t *testing.T) {
	result := encoder.ErrorToHtml{}.IsAble(errors.New("entity not found"))
	assert.True(t, result)
}

func TestNotCorrectErrorCanNotConvertToHtml(t *testing.T) {
	app := setUp()
	encoders := []inter.Encoder{encoder.ErrorToHtml{}}
	encoder := encoder.ErrorToHtml{}
	result, err := encoder.EncodeThrough(app, "foo", encoders)

	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can't convert object to html in error format")
}

func TestOneErrorToHtmlWithoutTemplate(t *testing.T) {
	app := setUp()
	app.Bind("config.App.Debug", false)
	encoder := encoder.ErrorToHtml{}
	result, err := encoder.EncodeThrough(app, errors.New("entity not found"), test.HtmlEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "Entity not found", result)
	assert.NotContains(t, result, "<p>")
	assert.NotContains(t, result, "errors_to_html.go")
}

func TestOneErrorToHtmlOnProduction(t *testing.T) {
	app := setUp()
	app.Bind("config.App.Debug", false)
	encoder := encoder.ErrorToHtml{View: test.NewViewErrorMock}
	result, err := encoder.EncodeThrough(app, errors.New("entity not found"), test.HtmlEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "<h1>500</h1>\n<h2>Entity not found</h2>\n", result)
	assert.NotContains(t, result, "<p>")
	assert.NotContains(t, result, "errors_to_html.go")
}

func TestOneErrorToHtmlOnDevelopmentWithStackTrace(t *testing.T) {
	app := setUp()
	app.Bind("config.App.Debug", true)
	encoder := encoder.ErrorToHtml{View: test.NewViewErrorMock}
	result, err := encoder.EncodeThrough(app, errors.New("entity not found"), test.HtmlEncoders)

	assert.Nil(t, err)
	assert.Contains(t, result, "Entity not found")
	assert.Contains(t, result, "errors_to_html_test.go")
}

//goland:noinspection GoNilness
func TestTemplateErrorToHtml(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorToHtml{View: test.NewViewErrorMock}.
		EncodeThrough(
			app,
			errors.New("entity not found"), test.HtmlEncoders,
		)

	assert.NoError(t, err)
	assert.Equal(t, "<h1>500</h1>\n<h2>Entity not found</h2>\n", result)
}

func TestSystemErrorConvertToHtml(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, errors.New("bad request"), []inter.Encoder{encoder.InterfaceToHtml{}})

	assert.Equal(t, "no encoder found to handle error: bad request", result)
	assert.EqualError(t, err, "no encoder found to handle error: bad request")
}
