package encoder

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
)

func RequestWithFormToValue(request inter.Request) support.Value {
	formMap, err := support.NewValueE(request.Source().Form)
	if err != nil {
		return support.NewValue(err)
	}
	return formMap
}
