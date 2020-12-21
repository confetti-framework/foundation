package encoder

import (
	"encoding/json"
	"github.com/confetti-framework/contract/inter"
)

type InterfaceToJson struct{}

func (j InterfaceToJson) IsAble(_ interface{}) bool {
	return true
}

func (j InterfaceToJson) EncodeThrough(_ inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	if object == nil {
		return "", nil
	}

	result, err := json.Marshal(object)
	return string(result), err
}
