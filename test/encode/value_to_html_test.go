package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_value_to_html_with_string_can_not_encode(t *testing.T) {
	require.False(t, encoder.RawToHtml{}.IsAble(""))
}

func Test_value_to_html_can_encode(t *testing.T) {
	require.True(t, encoder.RawToHtml{}.IsAble(support.NewValue("foo")))
}

func Test_value_to_hml_with_string(t *testing.T) {
	app := setUp()
	result, err := encoder.RawToHtml{}.EncodeThrough(app, "foo", nil)
	require.Equal(t, "", result)
	require.EqualError(t, err, "can not encode to html with an unsupported type string")
}

func Test_value_to_html_with_valid_value(t *testing.T) {
	app := setUp()
	result, err := encoder.RawToHtml{}.EncodeThrough(app, support.NewValue("foo"), mock.HtmlEncoders)
	require.NoError(t, err)
	require.Equal(t, "foo", result)
}
