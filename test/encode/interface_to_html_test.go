package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilCanTransformToHtml(t *testing.T) {
	assert.True(t, encoder.InterfaceToHtml{}.IsAble(nil))
}

func TestStringCanTransformToHtml(t *testing.T) {
	object := "foo"
	assert.True(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func TestIntCanTransformToHtml(t *testing.T) {
	object := 12
	assert.False(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func TestEmptyStructCanTransformToHtml(t *testing.T) {
	object := foo{}
	assert.False(t, encoder.InterfaceToHtml{}.IsAble(object))
}

func TestEncodeNilToHtml(t *testing.T) {
	result, err := encoder.InterfaceToHtml{}.EncodeThrough(nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, "", result)
}

func TestEncodeStringToHtml(t *testing.T) {
	result, err := encoder.InterfaceToHtml{}.EncodeThrough(nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, "", result)
}
