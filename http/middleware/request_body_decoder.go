package middleware

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/http/request_helper"
)

type RequestBodyDecoder struct{}

// This ensures the request can be decoded from JSON.
func (r RequestBodyDecoder) Handle(request inter.Request, next inter.Next) inter.Response {
	switch {
	case request_helper.IsJson(request):
		request.App().Singleton(inter.RequestBodyDecoder, encoder.JsonToValue)
	}

	return next(request)
}
