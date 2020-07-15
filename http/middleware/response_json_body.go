package middleware

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support/transformer"
)

type ResponseJsonBody struct{}

// This ensures that if the request contains a JSON header, the content can be decoded.
func (j ResponseJsonBody) Handle(request inter.Request, next inter.Next) inter.Response {
	request.App().Singleton("response_body_decoder", transformer.ValueToJson)

	return next(request)
}
