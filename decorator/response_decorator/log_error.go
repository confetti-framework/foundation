package response_decorator

import (
	"github.com/lanvard/contract/inter"
)

type LogError struct{}

func (l LogError) Decorate(response inter.Response) inter.Response {
	if err, ok := response.Content().(error); ok {
		response.App().Log().Emergency(err.Error())
	}

	return response
}
