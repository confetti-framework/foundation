package transformer

import (
	"github.com/lanvard/foundation/transformer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalStructIsNotValid(t *testing.T) {
	assert.False(t, transformer.JsonReaderToJson{}.Transformable(foo{}))
}

func TestStructWithJsonReaderIsValid(t *testing.T) {
	assert.True(t, transformer.JsonReaderToJson{}.Transformable(JsonReader{}))
}

func TestTransformNormalStructToJson(t *testing.T) {
	assert.PanicsWithValue(t, "can not transform to json with an unsupported type", func() {
		transformer.JsonReaderToJson{}.Transform(foo{})
	})
}

type JsonReader struct {
	data interface{}
}

func (j JsonReader) Json() interface{} {
	return j.data
}
