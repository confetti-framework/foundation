package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonResponseFromEmptyString(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "", test.JsonEncoders)

	assert.Equal(t, "\"\"", result)
}

func TestJsonResponseFromJsonString(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "{\"height\": 12}", test.JsonEncoders)

	assert.Equal(t, "{\"height\": 12}", result)
}

func TestStringToJson(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "foo", test.JsonEncoders)

	assert.Equal(t, "\"foo\"", result)
}
