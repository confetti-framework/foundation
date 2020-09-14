package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalStructIsNotValid(t *testing.T) {
	assert.False(t, encoder.JsonReaderToJson{}.IsAble(foo{}))
}

func TestStructWithJsonReaderIsValid(t *testing.T) {
	assert.True(t, encoder.JsonReaderToJson{}.IsAble(jsonReader{}))
}

func TestTransformNormalStructToJson(t *testing.T) {
	_, err := encoder.JsonReaderToJson{}.EncodeThrough(foo{}, nil)
	assert.EqualError(t, err, "can not encode to json with an unsupported type encode.foo")
}

func TestTransformJsonReaderWithStringToJson(t *testing.T) {
	data := jsonReader{map[string]string{"Unit": "gigatonne"}}
	result, err := encoder.JsonReaderToJson{}.EncodeThrough(data, outcome.JsonEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "{\"Unit\":\"gigatonne\"}", result)
}

func TestTransformJsonReaderWithValueToJson(t *testing.T) {
	value := support.NewValue(map[string]interface{}{"Unit": "megatonne"})

	result, err := encoder.JsonReaderToJson{}.EncodeThrough(jsonReader{value}, outcome.JsonEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "{\"Unit\":\"megatonne\"}", result)
}

func TestJsonReaderWithNilToJson(t *testing.T) {
	result, err := encoder.JsonReaderToJson{}.EncodeThrough(jsonReader{nil}, outcome.JsonEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "null", result)
}

type jsonReader struct {
	Data interface{}
}

func (j jsonReader) Json() interface{} {
	return j.Data
}
