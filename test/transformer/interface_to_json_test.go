package transformer

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/middleware"
	"github.com/lanvard/foundation/transformer"
	"github.com/lanvard/routing/outcome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonResponseWithoutResponseEncoder(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	response := middleware.ResponseJsonBody{}.Handle(request, func(request inter.Request) inter.Response {
		return outcome.Json(foo{12})
	})

	response.SetApp(request.App())
	response.App().Singleton(inter.ResponseBodyEncoder, nil)

	_, err := response.ContentE()
	assert.NotNil(t, err)
	assert.Equal(
		t,
		"no response encoder found",
		err.Error(),
	)
}

func TestNilCanTransformToJson(t *testing.T) {
	assert.True(t, transformer.InterfaceToJson{}.Transformable(nil))
}

func TestStringCanTransformToJson(t *testing.T) {
	object := "foo"
	assert.True(t, transformer.InterfaceToJson{}.Transformable(object))
}

func TestIntCanTransformToJson(t *testing.T) {
	object := 12
	assert.True(t, transformer.InterfaceToJson{}.Transformable(object))
}

func TestEmptyStructCanTransformToJson(t *testing.T) {
	object := foo{}
	assert.True(t, transformer.InterfaceToJson{}.Transformable(object))
}

func TestTransformStructWithIntToJson(t *testing.T) {
	object := foo{12}
	result, err := transformer.InterfaceToJson{}.Transform(object)
	assert.Nil(t, err)
	assert.Equal(t, "{\"height\":12}", result)
}

func TestTransformStructWithFloat(t *testing.T) {
	object := foo{12.34}
	result, err := transformer.InterfaceToJson{}.Transform(object)
	assert.Nil(t, err)
	assert.Equal(t, "{\"height\":12.34}", result)
}

func TestTransformNilToJson(t *testing.T) {
	result, err := transformer.InterfaceToJson{}.Transform(nil)
	assert.Nil(t, err)
	assert.Equal(t, "null", result)
}

type foo struct {
	Height float32 `json:"height"`
}
