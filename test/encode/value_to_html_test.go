package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValueToHtmlWithStringCanNotEncode(t *testing.T) {
	assert.False(t, encoder.RawToHtml{}.IsAble(""))
}

func TestValueToHtmlCanEncode(t *testing.T) {
	assert.True(t, encoder.RawToHtml{}.IsAble(support.NewValue("foo")))
}

func TestValueToHmlWithString(t *testing.T) {
	result, err := encoder.RawToHtml{}.EncodeThrough("foo", nil)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can not encode to html with an unsupported type string")
}

func TestValueToHtmlWithValidValue(t *testing.T) {
	result, err := encoder.RawToHtml{}.EncodeThrough(support.NewValue("foo"), outcome.HtmlEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "foo", result)
}
