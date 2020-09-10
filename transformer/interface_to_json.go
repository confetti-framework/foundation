package transformer

import (
	"encoding/json"
)

type InterfaceToJson struct{}

func (j InterfaceToJson) Transformable(_ interface{}) bool {
	return true
}

func (j InterfaceToJson) Transform(object interface{}) (string, error) {
	result, err := json.Marshal(object)
	return string(result), err
}
