package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/http/request_helper"
)

type RequestBodyDecoder struct{}

// This ensures the request can be decoded from JSON.
func (r RequestBodyDecoder) Handle(request inter.Request, next inter.Next) inter.Response {

	switch {
	case request_helper.HasJson(request):
		request.App().Singleton(inter.RequestBodyDecoder, encoder.RequestWithJsonToValue)
	case request_helper.HasMultiPartFormData(request):
		request.App().Singleton(inter.RequestBodyDecoder, encoder.RequestWithFormToValue)
	}

	return next(request)
}
