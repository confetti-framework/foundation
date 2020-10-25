package response_decorator

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
)

type LogError struct{}

func (l LogError) Decorate(response inter.Response) inter.Response {
	if err, ok := response.Content().(error); ok {
		level, _ := errors.FindLevel(err)
		response.App().Log().LogWith(level, err.Error(), err)
	}

	return response
}
