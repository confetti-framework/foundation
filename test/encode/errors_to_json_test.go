package encode

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/middleware"
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
	encoders := []inter.Encoder{encoder.ErrorToJson{}}
	result, err := encoder.ErrorToJson{}.EncodeThrough("foo", encoders)

	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can't convert object to json in error format")
}

func TestOneErrorToJson(t *testing.T) {
	result, err := encoder.ErrorToJson{}.EncodeThrough(errors.New("entity not found"), outcome.JsonEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"Entity not found\"}]}", result)
}

func TestOneErrorWithLongErrorMessage(t *testing.T) {
	result, err := encoder.ErrorToJson{}.EncodeThrough(
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

	assert.Nil(t, err)
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"This is a long error message, "+
		"this is a long error message, this is a long error message, this is a long error message, this is a long error message, this is a long error message, this is a long error message\"}]}", result)
}

func Test_request_without_content_type(t *testing.T) {
	// Given
	request := http.NewRequest(http.Options{
		App: foundation.NewApp(),
	})

	// When
	response := middleware.RequestBodyDecoder{}.Handle(request, func(request inter.Request) inter.Response {
		value := request.Body("data.foo.0.bar.1.bar")
		return outcome.Html(value.Error().Error())
	})
	response.SetApp(request.App())

	// Then
	assert.Equal(t, "{\"jsonapi\":{\"version\":\"1.0\"},\"errors\":[{\"title\":\"Content-Type not supported\"}]}", response.Content())
}

func TestSystemErrorConvertToJson(t *testing.T) {
	result, err := encoder.EncodeThrough("", []inter.Encoder{})

	assert.Equal(t, "No encoder found to handle error: no encoder found to encode response body with type string", result)
	assert.EqualError(t, err, "No encoder found to handle error: no encoder found to encode response body with type string")
}
