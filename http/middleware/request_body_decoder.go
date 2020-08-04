package middleware

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/helper"
	"github.com/lanvard/support/transformer"
)

type RequestBodyDecoder struct{}

// This ensures the request can be decoded from JSON.
func (j RequestBodyDecoder) Handle(request inter.Request, next inter.Next) inter.Response {
	switch {
	case helper.IsJson(request):
		request.App().Singleton(inter.RequestBodyDecoder, transformer.JsonToValue)
	}

	return next(request)
}
