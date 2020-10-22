package middleware

import (
	"github.com/lanvard/contract/inter"
)

type ResponseEncoder struct {
	Encoder func(content interface{}) inter.Response
}

func (r ResponseEncoder) Handle(request inter.Request, next inter.Next) inter.Response {
	request.App().Bind("response_encoder", r.Encoder)
	return next(request)
}
