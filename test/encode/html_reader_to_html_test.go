package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyHtmlReaderToHtmlCanNotEncode(t *testing.T) {
	assert.False(t, encoder.HtmlReaderToJson{}.IsAble(""))
}

func TestHtmlReaderCanEncode(t *testing.T) {
	assert.True(t, encoder.HtmlReaderToJson{}.IsAble(htmlReader{}))
}

func TestHtmlReaderToHmlWithString(t *testing.T) {
	result, err := encoder.HtmlReaderToJson{}.EncodeThrough("foo", nil)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can not transform to html with an unsupported type string")
}

func TestHtmlReaderToHtmlWithValidHtmlReader(t *testing.T) {
	result, err := encoder.HtmlReaderToJson{}.EncodeThrough(htmlReader{}, nil)
	assert.Nil(t, err)
	assert.Equal(t, "foo", result)
}

type htmlReader struct{}

func (h htmlReader) Html() string {
	return "foo"
}
