package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHtmlReaderToHtmlWithStringCanNotEncode(t *testing.T) {
	assert.False(t, encoder.HtmlReaderToHtml{}.IsAble(""))
}

func TestHtmlReaderToHtmlWithEmptyHtmlReaderCanEncode(t *testing.T) {
	assert.True(t, encoder.HtmlReaderToHtml{}.IsAble(htmlReader{}))
}

func TestHtmlReaderToHmlWithString(t *testing.T) {
	app := setUp()
	result, err := encoder.HtmlReaderToHtml{}.EncodeThrough(app, "foo", nil)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can not encode to html with an unsupported type string")
}

func TestHtmlReaderToHtmlWithValidHtmlReader(t *testing.T) {
	app := setUp()
	result, err := encoder.HtmlReaderToHtml{}.EncodeThrough(app, htmlReader{}, nil)
	assert.NoError(t, err)
	assert.Equal(t, "foo", result)
}

type htmlReader struct{}

func (h htmlReader) String() string {
	return "foo"
}
