package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyStringToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(""), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "\"\"", result)
}

func TestNilStringToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(nil), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "null", result)
}

func TestValueWithStringToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue("foo"), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "\"foo\"", result)
}

func TestBoolTrueToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(true), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "true", result)
}

func TestBoolFalseToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(false), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "false", result)
}

func TestPositiveNumberToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(100), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "100", result)
}

func TestNegativeNumberToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(-100), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "-100", result)
}

func TestFloatToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewValue(1.2), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "1.2", result)
}

func TestCollectionWithOneStringToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewCollection("foo"), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[\"foo\"]", result)
}

func TestCollectionWithThoStringsToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewCollection("foo", "bar"), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[\"foo\",\"bar\"]", result)
}

func TestCollectionWithThoNumbersToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewCollection(12, 14), defaultEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[12,14]", result)
}

func TestCollectionWithThoFloatToJson(t *testing.T) {
	result, err := encoder.EncodeThrough(support.NewCollection(1.5, 0.4), defaultEncoders)
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
	assert.EqualError(t, err, "can not transform to json with an unsupported type string")
}
