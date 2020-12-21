package response_decorator

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
)

type LogError struct{}

func (l LogError) Decorate(response inter.Response) inter.Response {
	if err, ok := response.GetContent().(error); ok {
		app := response.App()

		if l.ignore(app, err) {
			return response
		}

		level, _ := errors.FindLevel(err)
		app.Log().LogWith(level, err.Error(), err)
	}

	return response
}

func (l LogError) ignore(app inter.App, err error) bool {
	toIgnoreRaw := app.Make("config.Errors.NoLogging").([]interface{})
	for _, rawErr := range toIgnoreRaw {
		if errors.Is(err, rawErr.(error)) {
			return true
		}
	}

	return false
}
