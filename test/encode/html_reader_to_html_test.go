package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_stringer_to_html_with_string_can_not_encode(t *testing.T) {
	require.False(t, encoder.StringerToHtml{}.IsAble(""))
}

func Test_stringer_to_html_with_empty_stringer_can_encode(t *testing.T) {
	require.True(t, encoder.StringerToHtml{}.IsAble(content{}))
}

func Test_stringer_to_hml_with_string(t *testing.T) {
	app := setUp()
	result, err := encoder.StringerToHtml{}.EncodeThrough(app, "foo", nil)
	require.Equal(t, "", result)
	require.EqualError(t, err, "can not encode to html with an unsupported type string")
}

func Test_stringer_to_html_with_valid_stringer(t *testing.T) {
	app := setUp()
	result, err := encoder.StringerToHtml{}.EncodeThrough(app, content{}, nil)
	require.NoError(t, err)
	require.Equal(t, "foo", result)
}

type content struct{}

func (h content) String() string {
	return "foo"
}
