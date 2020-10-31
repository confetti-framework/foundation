package response_decorator

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
)

type LogError struct{}

func (l LogError) Decorate(response inter.Response) inter.Response {
	if err, ok := response.GetContent().(error); ok {
		app := response.App()
		errorsToIgnore := app.Make("config.Errors.NoLogging").([]error)
		if errors.Is(err, errorsToIgnore...) {
			return response
		}

		level, _ := errors.FindLevel(err)
		app.Log().LogWith(level, err.Error(), err)
	}

	return response
}
