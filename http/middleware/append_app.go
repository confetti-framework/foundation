package middleware

import (
	"github.com/confetti-framework/contract/inter"
)

type AppendApp struct{}

func (a AppendApp) Handle(request inter.Request, next inter.Next) inter.Response {
	response := next(request)
	if response.App() == nil {
		response.SetApp(request.App())
	}

	return response
}
