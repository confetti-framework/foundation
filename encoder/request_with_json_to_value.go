package encoder

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support"
	"io/ioutil"
)

func RequestWithJsonToValue(request inter.Request) support.Value {
	rawBody, err := ioutil.ReadAll(request.Source().Body)
	if err != nil {
		return support.NewValue(err)
	}

	return JsonToValue(string(rawBody))
}
