package response_decorator

import (
	"github.com/lanvard/contract/inter"
	net "net/http"
)

type HttpStatus struct{}

func (h HttpStatus) Decorate(response inter.Response) inter.Response {
	if _, ok := response.Content().(error); ok {
		httpStatus := net.StatusInternalServerError
		response.SetStatus(httpStatus)
	}

	return response
}
