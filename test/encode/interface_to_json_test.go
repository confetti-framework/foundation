package encode

import (
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonResponseWithoutResponseEncoder(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	response := outcome.NewResponse(outcome.Options{
		Content: foo{12},
	})

	response.SetApp(request.App())

	_, err := response.ContentE()
	assert.NotNil(t, err)
	assert.Equal(
		t,
		"no response encoder found",
		err.Error(),
	)
}

func TestNilCanTransformToJson(t *testing.T) {
	assert.True(t, encoder.InterfaceToJson{}.IsAble(nil))
}

func TestStringCanTransformToJson(t *testing.T) {
	object := "foo"
	assert.True(t, encoder.InterfaceToJson{}.IsAble(object))
}

func TestIntCanTransformToJson(t *testing.T) {
	object := 12
	assert.True(t, encoder.InterfaceToJson{}.IsAble(object))
}

func TestEmptyStructCanTransformToJson(t *testing.T) {
	object := foo{}
	assert.True(t, encoder.InterfaceToJson{}.IsAble(object))
}

func TestTransformStructWithIntToJson(t *testing.T) {
	app := setUp()
	object := foo{12}
	result, err := encoder.InterfaceToJson{}.EncodeThrough(app, object, nil)
	assert.Nil(t, err)
	assert.Equal(t, "{\"height\":12}", result)
}

func TestTransformStructWithFloat(t *testing.T) {
	app := setUp()
	object := foo{12.34}
	result, err := encoder.InterfaceToJson{}.EncodeThrough(app, object, nil)
	assert.Nil(t, err)
	assert.Equal(t, "{\"height\":12.34}", result)
}

func TestTransformNilToJson(t *testing.T) {
	app := setUp()
	result, err := encoder.InterfaceToJson{}.EncodeThrough(app, nil, nil)
	assert.Nil(t, err)
	assert.Equal(t, "", result)
}

type foo struct {
	Height float32 `json:"height"`
}
