package response_decorator

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	net "net/http"
)

type FilterSensitiveError struct{}

func (c FilterSensitiveError) Decorate(response inter.Response) inter.Response {
	if response.App().Make("config.App.Debug").(bool) {
		return response
	}

	// For security reasons system errors should not be exposed
	if _, ok := response.GetContent().(error); ok && isServerError(response) {
		response.Content(errors.New("an error has occurred"))
	}

	return response
}

func isServerError(response inter.Response) bool {
	return response.GetStatus() == 0 || response.GetStatus() >= net.StatusInternalServerError
}
