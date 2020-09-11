package middleware

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/encoder"
)

type ResponseJsonBody struct{}

// This ensures that if the request contains a JSON header, the content can be decoded.
func (j ResponseJsonBody) Handle(request inter.Request, next inter.Next) inter.Response {
	request.App().Singleton(inter.Encoders, []inter.Encoder{
		encoder.JsonReaderToJson{},
		encoder.RawToJson{},
		encoder.JsonToJson{},
		encoder.InterfaceToJson{},
	})

	return next(request)
}
