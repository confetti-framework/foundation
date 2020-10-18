package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyStringToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(""), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "\"\"", result)
}

func TestNilStringToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(nil), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "null", result)
}

func TestValueWithStringToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue("foo"), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "\"foo\"", result)
}

func TestBoolTrueToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(true), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "true", result)
}

func TestBoolFalseToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(false), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "false", result)
}

func TestPositiveNumberToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(100), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "100", result)
}

func TestNegativeNumberToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(-100), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "-100", result)
}

func TestFloatToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(1.2), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "1.2", result)
}

func TestCollectionWithOneStringToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewCollection("foo"), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[\"foo\"]", result)
}

func TestCollectionWithThoStringsToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewCollection("foo", "bar"), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[\"foo\",\"bar\"]", result)
}

func TestCollectionWithThoNumbersToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewCollection(12, 14), outcome.JsonEncoders)
	assert.Nil(t, err)
	assert.Equal(t, "[12,14]", result)
}

func TestCollectionWithThoFloatToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewCollection(1.5, 0.4), outcome.JsonEncoders)
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
	app := setUp()
	_, err := encoder.RawToJson{}.EncodeThrough(app, "foo", nil)
	assert.EqualError(t, err, "can not encode to json with an unsupported type string")
}
