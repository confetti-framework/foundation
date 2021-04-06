package encode

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/stretchr/testify/require"
	"html/template"
	"testing"
)

func Test_can_not_convert_string_to_json_error(t *testing.T) {
	result := encoder.ErrorsToJson{}.IsAble("Foo")
	require.False(t, result)
}

func Test_one_error_can_convert_to_json(t *testing.T) {
	result := encoder.ErrorsToJson{}.IsAble(errors.New("entity not found"))
	require.True(t, result)
}

func Test_slice_with_errors_can_convert_to_json(t *testing.T) {
	result := encoder.ErrorsToJson{}.IsAble([]error{})
	require.True(t, result)
}

func Test_not_correct_error_can_not_convert_to_json(t *testing.T) {
	app := setUp()
	encoders := []inter.Encoder{encoder.ErrorsToJson{}}
	result, err := encoder.ErrorsToJson{}.EncodeThrough(app, "foo", encoders)

	require.Equal(t, "", result)
	require.EqualError(t, err, "can't convert object to json in error format")
}

func Test_one_error_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorsToJson{}.EncodeThrough(app, errors.New("entity not found"), mock.JsonEncoders)

	require.NoError(t, err)
	require.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"Entity not found\"}]}", result)
}

func Test_one_error_with_long_error_message(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorsToJson{}.EncodeThrough(
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

	require.NoError(t, err)
	require.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"This is a long error message, "+
		"this is a long error message, this is a long error message, this is a long error message, this is a long "+
		"error message, this is a long error message, this is a long error message\"}]}", result)
}

func Test_system_error_convert_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, "", []inter.Encoder{})

	require.Equal(t, "no encoder found to handle error: no encoder found to encode response body with type string", result)
	require.EqualError(t, err, "no encoder found to handle error: no encoder found to encode response body with type string")
}

func Test_convert_slice_with_one_error_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorsToJson{}.EncodeThrough(
		app,
		[]error{errors.New("entity not found")},
		mock.JsonEncoders,
	)

	require.NoError(t, err)
	require.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"Entity not found\"}]}", result)
}

func Test_convert_slice_without_errors_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorsToJson{}.EncodeThrough(
		app,
		[]error{},
		mock.JsonEncoders,
	)

	require.NoError(t, err)
	require.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[]}", result)
}

func Test_convert_slice_with_two_errors_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorsToJson{}.EncodeThrough(
		app,
		[]error{errors.New("error one"), errors.New("error two")},
		mock.JsonEncoders,
	)

	require.NoError(t, err)
	require.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"Error one\"},{\"title\":\"Error two\"}]}", result)
}

func setUp() *foundation.Application {
	app := foundation.NewApp()
	app.Bind("config.App.Debug", false)
	app.Bind("template_builder", func(template *template.Template) (*template.Template, error) {
		return template.Parse("{{define \"footer\"}}\n    <footer>contact information</footer>{{end}}")
	})
	return app
}
