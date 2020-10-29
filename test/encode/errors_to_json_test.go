package encode

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanNotConvertStringToJsonError(t *testing.T) {
	result := encoder.ErrorToJson{}.IsAble("Foo")
	assert.False(t, result)
}

func TestOneErrorCanConvertToJson(t *testing.T) {
	result := encoder.ErrorToJson{}.IsAble(errors.New("entity not found"))
	assert.True(t, result)
}

func TestNotCorrectErrorCanNotConvertToJson(t *testing.T) {
	app := setUp()
	encoders := []inter.Encoder{encoder.ErrorToJson{}}
	result, err := encoder.ErrorToJson{}.EncodeThrough(app, "foo", encoders)

	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can't convert object to json in error format")
}

func TestOneErrorToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorToJson{}.EncodeThrough(app, errors.New("entity not found"), outcome.JsonEncoders)

	assert.NoError(t, err)
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"Entity not found\"}]}", result)
}

func TestOneErrorWithLongErrorMessage(t *testing.T) {
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
		outcome.JsonEncoders,
	)

	assert.NoError(t, err)
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"This is a long error message, "+
		"this is a long error message, this is a long error message, this is a long error message, this is a long "+
		"error message, this is a long error message, this is a long error message\"}]}", result)
}

func TestSystemErrorConvertToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, "", []inter.Encoder{})

	assert.Equal(t, "no encoder found to handle error: no encoder found to encode response body with type string", result)
	assert.EqualError(t, err, "no encoder found to handle error: no encoder found to encode response body with type string")
}

func setUp() *foundation.Application {
	return foundation.NewApp()
}
