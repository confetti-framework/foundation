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

func TestNilCanTransformToStruct(t *testing.T) {
	assert.True(t, transformer.StructToJson{}.CanTransform(nil))
}

func TestStringCanNotTransformToStruct(t *testing.T) {
	object := "foo"
	assert.False(t, transformer.StructToJson{}.CanTransform(object))
}

func TestIntCanNotTransform(t *testing.T) {
	object := 12
	assert.False(t, transformer.StructToJson{}.CanTransform(object))
}



type Foo struct {
	Height float32
}
