package response_decorator

import (
	"github.com/lanvard/contract/inter"
	"github.com/pkg/errors"
)

type FilterSensitiveError struct{}

func (c FilterSensitiveError) Decorate(response inter.Response) inter.Response {
	if response.App().Make("config.App.Debug").(bool) {
		return response
	}

	// For security reasons system exceptions should not be exposed
	content := response.Content()
	if _, ok := content.(error); ok {
		content = errors.New("an error has occurred")
	}

	response.SetContent(content)

	return response
}
