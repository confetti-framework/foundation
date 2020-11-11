package encode

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/test/mock"
	"github.com/stretchr/testify/assert"
	"html/template"
	"testing"
)

func Test_can_not_convert_string_to_json_error(t *testing.T) {
	result := encoder.ErrorToJson{}.IsAble("Foo")
	assert.False(t, result)
}

func Test_one_error_can_convert_to_json(t *testing.T) {
	result := encoder.ErrorToJson{}.IsAble(errors.New("entity not found"))
	assert.True(t, result)
}

func Test_not_correct_error_can_not_convert_to_json(t *testing.T) {
	app := setUp()
	encoders := []inter.Encoder{encoder.ErrorToJson{}}
	result, err := encoder.ErrorToJson{}.EncodeThrough(app, "foo", encoders)

	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can't convert object to json in error format")
}

func Test_one_error_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorToJson{}.EncodeThrough(app, errors.New("entity not found"), mock.JsonEncoders)

	assert.NoError(t, err)
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"Entity not found\"}]}", result)
}

func Test_one_error_with_long_error_message(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorToJson{}.EncodeThrough(
		app,
		errors.New(
			"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message, "+
				"this is a long error message",
		),
		mock.JsonEncoders,
	)

	assert.NoError(t, err)
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"This is a long error message, "+
		"this is a long error message, this is a long error message, this is a long error message, this is a long "+
		"error message, this is a long error message, this is a long error message\"}]}", result)
}

func Test_system_error_convert_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, "", []inter.Encoder{})

	assert.Equal(t, "no encoder found to handle error: no encoder found to encode response body with type string", result)
	assert.EqualError(t, err, "no encoder found to handle error: no encoder found to encode response body with type string")
}

func setUp() *foundation.Application {
	app := foundation.NewApp()
	app.Bind("config.App.Debug", false)
	app.Singleton("template_builder", func(template *template.Template) (*template.Template, error) {
		return template.ParseGlob(mock.TemplateByName("") + "/[^sub]*")
	})
	return app
}
