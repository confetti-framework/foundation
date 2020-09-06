package transformer

import (
	"encoding/json"
)

type InterfaceToJson struct{}

func (j InterfaceToJson) Transformable(object interface{}) bool {
	return true
}

func (j InterfaceToJson) Transform(object interface{}) string {
	result, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	return string(result)
}
