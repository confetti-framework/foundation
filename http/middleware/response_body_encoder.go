package middleware

import (
	"github.com/lanvard/contract/inter"
)

type ResponseJsonBody struct{}

// This ensures that if the request contains a JSON header, the content can be decoded.
func (j ResponseJsonBody) Handle(request inter.Request, next inter.Next) inter.Response {
	// request.App().Singleton(inter.ResponseBodyEncoder, transformer.StructToJson)
	return next(request)
}
