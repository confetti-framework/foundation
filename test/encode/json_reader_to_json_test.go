package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_normal_struct_is_not_valid(t *testing.T) {
	assert.False(t, encoder.JsonReaderToJson{}.IsAble(foo{}))
}

func Test_struct_with_json_reader_is_valid(t *testing.T) {
	assert.True(t, encoder.JsonReaderToJson{}.IsAble(jsonReader{}))
}

func Test_transform_normal_struct_to_json(t *testing.T) {
	app := setUp()
	_, err := encoder.JsonReaderToJson{}.EncodeThrough(app, foo{}, nil)
	assert.EqualError(t, err, "can not encode to json with an unsupported type encode.foo")
}

func Test_transform_json_reader_with_string_to_json(t *testing.T) {
	app := setUp()
	data := jsonReader{map[string]string{"Unit": "gigatonne"}}
	result, err := encoder.JsonReaderToJson{}.EncodeThrough(app, data, mock.JsonEncoders)

	assert.NoError(t, err)
	assert.Equal(t, "{\"Unit\":\"gigatonne\"}", result)
}

func Test_transform_json_reader_with_value_to_json(t *testing.T) {
	app := setUp()
	value := support.NewValue(map[string]interface{}{"Unit": "megatonne"})

	result, err := encoder.JsonReaderToJson{}.EncodeThrough(app, jsonReader{value}, mock.JsonEncoders)

	assert.NoError(t, err)
	assert.Equal(t, "{\"Unit\":\"megatonne\"}", result)
}

func Test_json_reader_with_nil_to_json(t *testing.T) {
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
