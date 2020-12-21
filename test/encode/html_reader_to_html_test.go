package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_stringer_to_html_with_string_can_not_encode(t *testing.T) {
	assert.False(t, encoder.StringerToHtml{}.IsAble(""))
}

func Test_stringer_to_html_with_empty_stringer_can_encode(t *testing.T) {
	assert.True(t, encoder.StringerToHtml{}.IsAble(content{}))
}

func Test_stringer_to_hml_with_string(t *testing.T) {
	app := setUp()
	result, err := encoder.StringerToHtml{}.EncodeThrough(app, "foo", nil)
	assert.Equal(t, "", result)
	assert.EqualError(t, err, "can not encode to html with an unsupported type string")
}

func Test_stringer_to_html_with_valid_stringer(t *testing.T) {
	app := setUp()
	result, err := encoder.StringerToHtml{}.EncodeThrough(app, content{}, nil)
	assert.NoError(t, err)
	assert.Equal(t, "foo", result)
}

type content struct{}

func (h content) String() string {
	return "foo"
}
