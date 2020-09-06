package transformer

import (
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
	assert.PanicsWithValue(t, "can not transform to json with an unsupported type transformer.foo", func() {
		transformer.JsonReaderToJson{}.Transform(foo{})
	})
}

func TestTransformJsonReaderWithStringToJson(t *testing.T) {
	data := jsonReader{map[string]string{"Unit": "gigatonne"}}
	result := transformer.JsonReaderToJson{}.Transform(data)

	assert.Equal(t, "{\"Unit\":\"gigatonne\"}", result)
}

func TestTransformJsonReaderWithValueToJson(t *testing.T) {
	value := support.NewValue(map[string]interface{}{"Unit": "megatonne"})
	result := transformer.JsonReaderToJson{}.Transform(jsonReader{value})

	assert.Equal(t, "{\"Unit\":\"megatonne\"}", result)
}

type jsonReader struct {
	Data interface{}
}

func (j jsonReader) Json() interface{} {
	return j.Data
}
