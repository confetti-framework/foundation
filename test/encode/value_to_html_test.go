package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/test"
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
	app := setUp()
	result, err := encoder.RawToHtml{}.EncodeThrough(app, "foo", nil)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can not encode to html with an unsupported type string")
}

func TestValueToHtmlWithValidValue(t *testing.T) {
	app := setUp()
	result, err := encoder.RawToHtml{}.EncodeThrough(app, support.NewValue("foo"), test.HtmlEncoders)
	assert.NoError(t, err)
	assert.Equal(t, "foo", result)
}
