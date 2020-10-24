package encode

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
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
	app := setUp()
	encoders := []inter.Encoder{encoder.ErrorToHtml{}}
	result, err := encoder.ErrorToHtml{}.EncodeThrough(app, "foo", encoders)

	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can't convert object to html in error format")
}

func TestOneErrorToHtml(t *testing.T) {
	app := setUp()
	result, err := encoder.ErrorToHtml{}.EncodeThrough(app, errors.New("entity not found"), outcome.HtmlEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "Entity not found", result)
}

func TestSystemErrorConvertToHtml(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, errors.New("bad request"), []inter.Encoder{encoder.InterfaceToHtml{}})

	assert.Equal(t, "no encoder found to handle error: bad request", result)
	assert.EqualError(t, err, "no encoder found to handle error: bad request")
}
