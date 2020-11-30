package encoder

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
)

func RequestWithFormToValue(request inter.Request) support.Value {
	source := request.Source()
	if err := source.ParseForm(); err != nil {
		return support.NewValue(err)
	}
	formMap, err := support.NewValueE(source.Form)
	if err != nil {
		return support.NewValue(err)
	}
	return formMap
}
