package response_decorator

import (
	"github.com/lanvard/contract/inter"
)

var ResponseDecorators = []inter.ResponseDecorator{
	FilterSensitiveError{},
}
