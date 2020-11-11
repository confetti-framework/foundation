package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/test/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_json_response_from_empty_string(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "", mock.JsonEncoders)

	assert.Equal(t, "\"\"", result)
}

func Test_json_response_from_json_string(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "{\"height\": 12}", mock.JsonEncoders)

	assert.Equal(t, "{\"height\": 12}", result)
}

func Test_string_to_json(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "foo", mock.JsonEncoders)

	assert.Equal(t, "\"foo\"", result)
}
