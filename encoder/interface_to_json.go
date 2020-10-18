package encoder

import (
	"encoding/json"
	"github.com/lanvard/contract/inter"
)

type InterfaceToJson struct{}

func (j InterfaceToJson) IsAble(_ interface{}) bool {
	return true
}

func (j InterfaceToJson) EncodeThrough(app inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	if object == nil {
		return "", nil
	}

	result, err := json.Marshal(object)
	return string(result), err
}
