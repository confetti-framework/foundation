package encode

import (
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/test/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_json_response_from_empty_string(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "", mock.JsonEncoders)

	require.Equal(t, "\"\"", result)
}

func Test_json_response_from_json_string(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "{\"height\": 12}", mock.JsonEncoders)

	require.Equal(t, "{\"height\": 12}", result)
}

func Test_string_to_json(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "foo", mock.JsonEncoders)

	require.Equal(t, "\"foo\"", result)
}
