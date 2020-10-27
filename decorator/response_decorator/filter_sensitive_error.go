package response_decorator

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
)

type FilterSensitiveError struct{}

func (c FilterSensitiveError) Decorate(response inter.Response) inter.Response {
	if response.App().Make("config.App.Debug").(bool) {
		return response
	}

	// For security reasons system errors should not be exposed
	if err, ok := response.Content().(error); ok {
		err = errors.New("an error has occurred")
		response.SetContent(err)
	}

	return response
}
