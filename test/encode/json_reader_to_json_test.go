package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/test/mock"
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
	app := setUp()
	_, err := encoder.JsonReaderToJson{}.EncodeThrough(app, foo{}, nil)
	assert.EqualError(t, err, "can not encode to json with an unsupported type encode.foo")
}

func TestTransformJsonReaderWithStringToJson(t *testing.T) {
	app := setUp()
	data := jsonReader{map[string]string{"Unit": "gigatonne"}}
	result, err := encoder.JsonReaderToJson{}.EncodeThrough(app, data, mock.JsonEncoders)

	assert.NoError(t, err)
	assert.Equal(t, "{\"Unit\":\"gigatonne\"}", result)
}

func TestTransformJsonReaderWithValueToJson(t *testing.T) {
	app := setUp()
	value := support.NewValue(map[string]interface{}{"Unit": "megatonne"})

	result, err := encoder.JsonReaderToJson{}.EncodeThrough(app, jsonReader{value}, mock.JsonEncoders)

	assert.NoError(t, err)
	assert.Equal(t, "{\"Unit\":\"megatonne\"}", result)
}

func TestJsonReaderWithNilToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.JsonReaderToJson{}.EncodeThrough(app, jsonReader{nil}, mock.JsonEncoders)

	assert.NoError(t, err)
	assert.Equal(t, "null", result)
}

type jsonReader struct {
	Data interface{}
}

func (j jsonReader) Json() interface{} {
	return j.Data
}
