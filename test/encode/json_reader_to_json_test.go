package encode

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/encoder"
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
	result, err := encoder.JsonReaderToJson{}.EncodeThrough(data, jsonEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "{\"Unit\":\"gigatonne\"}", result)
}

func TestTransformJsonReaderWithValueToJson(t *testing.T) {
	value := support.NewValue(map[string]interface{}{"Unit": "megatonne"})

	result, err := encoder.JsonReaderToJson{}.EncodeThrough(jsonReader{value}, jsonEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "{\"Unit\":\"megatonne\"}", result)
}

func TestJsonReaderWithNilToJson(t *testing.T) {
	result, err := encoder.JsonReaderToJson{}.EncodeThrough(jsonReader{nil}, jsonEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "null", result)
}

var jsonEncoders = []inter.Encoder{
	encoder.JsonReaderToJson{},
	encoder.RawToJson{},
	encoder.JsonToJson{},
	encoder.ErrorToJson{},
	encoder.InterfaceToJson{},
}

type jsonReader struct {
	Data interface{}
}

func (j jsonReader) Json() interface{} {
	return j.Data
}
