package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_nil_can_transform_to_html(t *testing.T) {
	require.True(t, encoder.InterfaceToHtml{}.IsAble(nil))
}

func Test_string_can_transform_to_html(t *testing.T) {
	object := "foo"
	require.True(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func Test_int_can_transform_to_html(t *testing.T) {
	object := 12
	require.True(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func Test_bool_can_transform_to_html(t *testing.T) {
	object := true
	require.True(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func Test_empty_struct_can_transform_to_html(t *testing.T) {
	object := foo{}
	require.False(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func Test_encode_nil_to_html(t *testing.T) {
	app := setUp()
	result, err := encoder.InterfaceToHtml{}.EncodeThrough(app, nil, nil)
	require.NoError(t, err)
	require.Equal(t, "", result)
}

func Test_encode_string_to_html(t *testing.T) {
	app := setUp()
	result, err := encoder.InterfaceToHtml{}.EncodeThrough(app, "The dog", nil)
	require.NoError(t, err)
	require.Equal(t, "The dog", result)
}

func Test_encode_bool_to_html(t *testing.T) {
	app := setUp()
	result, err := encoder.InterfaceToHtml{}.EncodeThrough(app, true, nil)
	require.NoError(t, err)
	require.Equal(t, "true", result)
}
