package encoder

import (
	"github.com/confetti-framework/support"
	"github.com/tidwall/gjson"
)

func JsonToValue(json string) support.Value {
	result := gjson.Parse(json).Value()

	return support.NewValue(result)
}
