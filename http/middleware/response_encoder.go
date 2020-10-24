package middleware

import (
	"github.com/lanvard/contract/inter"
)

type DefaultResponseEncoder struct {
	Encoder func(content interface{}) inter.Response
}

func (r DefaultResponseEncoder) Handle(request inter.Request, next inter.Next) inter.Response {
	request.App().Bind("default_response_encoder", r.Encoder)
	return next(request)
}
