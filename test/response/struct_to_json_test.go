package response

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http"
	"github.com/lanvard/foundation/http/middleware"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/support/transformer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonResponseWithoutResponseEncoder(t *testing.T) {
	options := http.Options{App: foundation.NewApp()}
	request := http.NewRequest(options)
	response := middleware.ResponseJsonBody{}.Handle(request, func(request inter.Request) inter.Response {
		return outcome.Json(Foo{12})
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
	assert.True(t, transformer.StructToJson{}.IsValid(nil))
}

func TestStringCanNotTransformToJson(t *testing.T) {
	object := "foo"
	assert.False(t, transformer.StructToJson{}.IsValid(object))
}

func TestIntCanNotTransformToJson(t *testing.T) {
	object := 12
	assert.False(t, transformer.StructToJson{}.IsValid(object))
}

func TestEmptyStructCanTransformToJson(t *testing.T) {
	object := Foo{}
	assert.True(t, transformer.StructToJson{}.IsValid(object))
}

func TestTransformUnsupportedTypeShouldPanic(t *testing.T) {
	assert.PanicsWithValue(t, "can not transform to json with an unsupported type", func() {
		transformer.StructToJson{}.Transform("foo")
	})
}

func TestTransformStructWithIntToJson(t *testing.T) {
	object := Foo{12}
	assert.Equal(t, "{\"height\":12}", transformer.StructToJson{}.Transform(object))
}

func TestTransformStructWithFloat(t *testing.T) {
	object := Foo{12.34}
	assert.Equal(t, "{\"height\":12.34}", transformer.StructToJson{}.Transform(object))
}

type Foo struct {
	Height float32 `json:"height"`
}
