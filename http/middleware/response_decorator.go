package middleware

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation/decorator/response_decorator"
)

type DecorateResponse struct{}

func (r DecorateResponse) Handle(request inter.Request, next inter.Next) inter.Response {
	response := next(request)
	decorators := response.App().Make("response_decorators").([]inter.ResponseDecorator)
	return response_decorator.Handler{Decorators: decorators}.Decorate(response)
}
