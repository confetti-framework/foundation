package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_empty_string_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(""), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "\"\"", result)
}

func Test_nil_string_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(nil), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "null", result)
}

func Test_value_with_string_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue("foo"), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "\"foo\"", result)
}

func Test_bool_true_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(true), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "true", result)
}

func Test_bool_false_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(false), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "false", result)
}

func Test_positive_number_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(100), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "100", result)
}

func Test_negative_number_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(-100), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "-100", result)
}

func Test_float_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewValue(1.2), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "1.2", result)
}

func Test_collection_with_one_string_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewCollection("foo"), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "[\"foo\"]", result)
}

func Test_merge_with_tho_strings_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewCollection("foo", "bar"), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "[\"foo\",\"bar\"]", result)
}

func Test_collection_with_tho_numbers_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewCollection(12, 14), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "[12,14]", result)
}

func Test_collection_with_tho_float_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.EncodeThrough(app, support.NewCollection(1.5, 0.4), mock.JsonEncoders)
	require.NoError(t, err)
	require.Equal(t, "[1.5,0.4]", result)
}

func Test_string_can_not_transform_value_expected(t *testing.T) {
	require.False(t, encoder.RawToJson{}.IsAble("foo"))
}

func Test_can_transform_value(t *testing.T) {
	require.True(t, encoder.RawToJson{}.IsAble(support.NewValue(foo{})))
}

func Test_transform_string_value_expected(t *testing.T) {
	app := setUp()
	_, err := encoder.RawToJson{}.EncodeThrough(app, "foo", nil)
	require.EqualError(t, err, "can not encode to json with an unsupported type string")
}
