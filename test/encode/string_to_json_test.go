package encode

import (
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonResponseFromEmptyString(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "", outcome.JsonEncoders)

	assert.Equal(t, "\"\"", result)
}

func TestJsonResponseFromJsonString(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "{\"height\": 12}", outcome.JsonEncoders)

	assert.Equal(t, "{\"height\": 12}", result)
}

func TestStringToJson(t *testing.T) {
	app := setUp()
	result, _ := encoder.EncodeThrough(app, "foo", outcome.JsonEncoders)

	assert.Equal(t, "\"foo\"", result)
}
