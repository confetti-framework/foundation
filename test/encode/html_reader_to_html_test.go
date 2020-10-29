package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringerToHtmlWithStringCanNotEncode(t *testing.T) {
	assert.False(t, encoder.StringerToHtml{}.IsAble(""))
}

func TestStringerToHtmlWithEmptyStringerCanEncode(t *testing.T) {
	assert.True(t, encoder.StringerToHtml{}.IsAble(content{}))
}

func TestStringerToHmlWithString(t *testing.T) {
	app := setUp()
	result, err := encoder.StringerToHtml{}.EncodeThrough(app, "foo", nil)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can not encode to html with an unsupported type string")
}

func TestStringerToHtmlWithValidStringer(t *testing.T) {
	app := setUp()
	result, err := encoder.StringerToHtml{}.EncodeThrough(app, content{}, nil)
	assert.NoError(t, err)
	assert.Equal(t, "foo", result)
}

type content struct{}

func (h content) String() string {
	return "foo"
}
