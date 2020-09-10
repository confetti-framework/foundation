package transformer

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/transformer"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalStructIsNotValid(t *testing.T) {
	assert.False(t, transformer.JsonReaderToJson{}.Transformable(foo{}))
}

func TestStructWithJsonReaderIsValid(t *testing.T) {
	assert.True(t, transformer.JsonReaderToJson{}.Transformable(jsonReader{}))
}

func TestTransformNormalStructToJson(t *testing.T) {
	_, err := transformer.JsonReaderToJson{}.TransformThrough(foo{}, nil)
	assert.EqualError(t, err, "can not transform to json with an unsupported type transformer.foo")
}

func TestTransformJsonReaderWithStringToJson(t *testing.T) {
	data := jsonReader{map[string]string{"Unit": "gigatonne"}}
	result, err := transformer.JsonReaderToJson{}.TransformThrough(data, defaultEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "{\"Unit\":\"gigatonne\"}", result)
}

func TestTransformJsonReaderWithValueToJson(t *testing.T) {
	value := support.NewValue(map[string]interface{}{"Unit": "megatonne"})

	result, err := transformer.TransformThrough(jsonReader{value}, defaultEncoders)

	assert.Nil(t, err)
	assert.Equal(t, "{\"Unit\":\"megatonne\"}", result)
}

var defaultEncoders = []inter.ResponseEncoder{
	transformer.JsonReaderToJson{},
	transformer.ValueToJson{},
	transformer.InterfaceToJson{},
}

type jsonReader struct {
	Data interface{}
}

func (j jsonReader) Json() interface{} {
	return j.Data
}
