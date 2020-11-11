package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/test/mock"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_value_to_html_with_string_can_not_encode(t *testing.T) {
	assert.False(t, encoder.RawToHtml{}.IsAble(""))
}

func Test_value_to_html_can_encode(t *testing.T) {
	assert.True(t, encoder.RawToHtml{}.IsAble(support.NewValue("foo")))
}

func Test_value_to_hml_with_string(t *testing.T) {
	app := setUp()
	result, err := encoder.RawToHtml{}.EncodeThrough(app, "foo", nil)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can not encode to html with an unsupported type string")
}

func Test_value_to_html_with_valid_value(t *testing.T) {
	app := setUp()
	result, err := encoder.RawToHtml{}.EncodeThrough(app, support.NewValue("foo"), mock.HtmlEncoders)
	assert.NoError(t, err)
	assert.Equal(t, "foo", result)
}
