package encoder

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support"
)

func RequestWithFormToValue(request inter.Request) support.Value {
	source := request.Source()
	if err := source.ParseForm(); err != nil {
		return support.NewValue(err)
	}
	formMap := support.NewValue(source.Form)
	if formMap.Error() != nil {
		return support.NewValue(formMap.Error())
	}
	return formMap
}
