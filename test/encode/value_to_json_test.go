package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyStringToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(""), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "\"\"", result)
}

func TestNilStringToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(nil), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "null", result)
}

func TestValueWithStringToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue("foo"), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "\"foo\"", result)
}

func TestBoolTrueToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(true), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "true", result)
}

func TestBoolFalseToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(false), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "false", result)
}

func TestPositiveNumberToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(100), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "100", result)
}

func TestNegativeNumberToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(-100), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "-100", result)
}

func TestFloatToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(1.2), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "1.2", result)
}

func TestCollectionWithOneStringToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewCollection("foo"), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[\"foo\"]", result)
}

func TestCollectionWithThoStringsToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewCollection("foo", "bar"), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[\"foo\",\"bar\"]", result)
}

func TestCollectionWithThoNumbersToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewCollection(12, 14), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[12,14]", result)
}

func TestCollectionWithThoFloatToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewCollection(1.5, 0.4), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[1.5,0.4]", result)
}

func TestStringCanNotTransformValueExpected(t *testing.T) {
	assert.False(t, encoder.RawToJson{}.IsAble("foo"))
}

func TestCanTransformValue(t *testing.T) {
	assert.True(t, encoder.RawToJson{}.IsAble(support.NewValue(foo{})))
}

func TestTransformStringValueExpected(t *testing.T) {
	_, err := encoder.RawToJson{}.EncodeThrough("foo", nil)
	assert.EqualError(t, err, "can not encode to json with an unsupported type string")
}
