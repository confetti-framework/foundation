package encode

import (
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonResponseFromEmptyString(t *testing.T) {
	response := outcome.NewResponse(outcome.Options{
		Content:  "",
		Encoders: outcome.JsonEncoders,
	})
	response.SetApp(foundation.NewApp())

	assert.Equal(t, "", response.Content())
}

func TestJsonResponseFromJsonString(t *testing.T) {
	response := outcome.NewResponse(outcome.Options{
		Content:  "{\"height\": 12}",
		Encoders: outcome.JsonEncoders,
	})
	response.SetApp(foundation.NewApp())

	assert.Equal(t, "{\"height\": 12}", response.Content())
}

func TestStringToJson(t *testing.T) {
	result, _ := encoder.EncodeThrough("foo", outcome.JsonEncoders)

	assert.Equal(t, "\"foo\"", result)
}
