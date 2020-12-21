package response_decorator

import "github.com/confetti-framework/contract/inter"

type Handler struct {
	Decorators []inter.ResponseDecorator
}

func (r Handler) Decorate(response inter.Response) inter.Response {
	for _, decorator := range r.Decorators {
		response = decorator.Decorate(response)
	}

	return response
}
