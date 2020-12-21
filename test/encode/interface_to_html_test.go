package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nil_can_transform_to_html(t *testing.T) {
	assert.True(t, encoder.InterfaceToHtml{}.IsAble(nil))
}

func Test_string_can_transform_to_html(t *testing.T) {
	object := "foo"
	assert.True(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func Test_int_can_transform_to_html(t *testing.T) {
	object := 12
	assert.False(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func Test_empty_struct_can_transform_to_html(t *testing.T) {
	object := foo{}
	assert.False(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func Test_encode_nil_to_html(t *testing.T) {
	app := setUp()
	result, err := encoder.InterfaceToHtml{}.EncodeThrough(app, nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, "", result)
}

func Test_encode_string_to_html(t *testing.T) {
	app := setUp()
	result, err := encoder.InterfaceToHtml{}.EncodeThrough(app, nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, "", result)
}
