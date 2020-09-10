package transformer

import (
	"encoding/json"
	"github.com/lanvard/contract/inter"
)

type InterfaceToJson struct{}

func (j InterfaceToJson) Transformable(_ interface{}) bool {
	return true
}

func (j InterfaceToJson) TransformThrough(object interface{}, _ []inter.ResponseEncoder) (string, error) {
	result, err := json.Marshal(object)
	return string(result), err
}
