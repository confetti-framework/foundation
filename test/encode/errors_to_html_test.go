package encode

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/routing/outcome"
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
	encoders := []inter.Encoder{encoder.ErrorToHtml{}}
	result, err := encoder.ErrorToHtml{}.EncodeThrough("foo", encoders)

	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can't convert object to html in error format")
}

func TestOneErrorToHtml(t *testing.T) {
	result, err := encoder.ErrorToHtml{}.EncodeThrough(errors.New("entity not found"), outcome.HtmlEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "Entity not found", result)
}

func TestSystemErrorConvertToHtml(t *testing.T) {
	result, err := encoder.EncodeThrough(errors.New("bad request"), []inter.Encoder{encoder.InterfaceToHtml{}})

	assert.Equal(t, "no encoder found to handle error: bad request", result)
	assert.EqualError(t, err, "no encoder found to handle error: bad request")
}
