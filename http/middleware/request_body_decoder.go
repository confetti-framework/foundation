package middleware

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/encoder"
	"github.com/lanvard/foundation/http/request_helper"
	"github.com/lanvard/routing/outcome"
)

type RequestBodyDecoder struct{}

// This ensures the request can be decoded from JSON.
func (j RequestBodyDecoder) Handle(request inter.Request, next inter.Next) inter.Response {
	switch {
	case request_helper.IsJson(request):
		request.App().Singleton(inter.RequestBodyDecoder, encoder.JsonToValue)
	default:
		return outcome.Json(errors.New("Content-Type not supported"))
	}

	return next(request)
}
