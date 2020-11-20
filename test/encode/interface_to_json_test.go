package encode

import (
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_json_response_without_response_encoder(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	response := outcome.NewResponse(outcome.Options{
		Content: foo{12},
	})

	response.SetApp(request.App())

	_, err := response.GetBodyE()
	assert.NotNil(t, err)
	assert.Equal(
		t,
		"get instance from container: key 'default_response_outcome': can not found value in map",
		err.Error(),
	)
}

func Test_nil_can_transform_to_json(t *testing.T) {
	assert.True(t, encoder.InterfaceToJson{}.IsAble(nil))
}

func Test_string_can_transform_to_json(t *testing.T) {
	object := "foo"
	assert.True(t, encoder.InterfaceToJson{}.IsAble(object))
}

func Test_int_can_transform_to_json(t *testing.T) {
	object := 12
	assert.True(t, encoder.InterfaceToJson{}.IsAble(object))
}

func Test_empty_struct_can_transform_to_json(t *testing.T) {
	object := foo{}
	assert.True(t, encoder.InterfaceToJson{}.IsAble(object))
}

func Test_transform_struct_with_int_to_json(t *testing.T) {
	app := setUp()
	object := foo{12}
	result, err := encoder.InterfaceToJson{}.EncodeThrough(app, object, nil)
	assert.NoError(t, err)
	assert.Equal(t, "{\"height\":12}", result)
}

func Test_transform_struct_with_float(t *testing.T) {
	app := setUp()
	object := foo{12.34}
	result, err := encoder.InterfaceToJson{}.EncodeThrough(app, object, nil)
	assert.NoError(t, err)
	assert.Equal(t, "{\"height\":12.34}", result)
}

func Test_transform_nil_to_json(t *testing.T) {
	app := setUp()
	result, err := encoder.InterfaceToJson{}.EncodeThrough(app, nil, nil)
	assert.NoError(t, err)
	assert.Equal(t, "", result)
}

type foo struct {
	Height float32 `json:"height"`
}
