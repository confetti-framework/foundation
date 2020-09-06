package middleware

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/http/request_helper"
	"github.com/lanvard/routing/outcome"
	"github.com/lanvard/support/transformer"
)

type RequestBodyDecoder struct{}

// This ensures the request can be decoded from JSON.
func (j RequestBodyDecoder) Handle(request inter.Request, next inter.Next) inter.Response {
	switch {
	case request_helper.IsJson(request):
		request.App().Singleton(inter.RequestBodyDecoder, transformer.JsonToValue)
	default:
		return outcome.Error(errors.New("Content-Type not supported"))
	}

	return next(request)
}
