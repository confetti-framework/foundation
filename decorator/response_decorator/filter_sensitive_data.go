package response_decorator

import (
	"github.com/lanvard/contract/inter"
)

type FilterSensitiveData struct{}

func (c FilterSensitiveData) Decorate(response inter.Response) inter.Response {
	return response
}
