package middleware

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/decorator/response_decorator"
)

type DecorateResponse struct{}

func (r DecorateResponse) Handle(request inter.Request, next inter.Next) inter.Response {
	response := next(request)
	decorators := request.App().Make("response_decorators").([]inter.ResponseDecorator)
	return response_decorator.Handler{Decorators: decorators}.Decorate(response)
}
