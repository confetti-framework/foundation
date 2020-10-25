package response_decorator

import (
	"github.com/lanvard/contract/inter"
	net "net/http"
)

var ResponseDecorators = []inter.ResponseDecorator{
	LogError{},
	FilterSensitiveError{},
	HttpStatus{ErrorDefault: net.StatusInternalServerError},
}
