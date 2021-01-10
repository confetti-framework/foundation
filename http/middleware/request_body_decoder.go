package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/encoder"
	"github.com/confetti-framework/foundation/http/http_helper"
)

type RequestBodyDecoder struct{}

// This ensures the request can be decoded from JSON and Multipart Form.
// Feel free to add your own decoders to support other formats as well.
func (r RequestBodyDecoder) Handle(request inter.Request, next inter.Next) inter.Response {

	switch {
	case http_helper.HasJson(request):
		request.App().Bind(inter.RequestBodyDecoder, encoder.RequestWithJsonToValue)
	case http_helper.HasMultiPartFormData(request):
		request.App().Bind(inter.RequestBodyDecoder, encoder.RequestWithFormToValue)
	}

	return next(request)
}
