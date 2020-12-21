package encoder

import (
	"encoding/json"
	"github.com/confetti-framework/contract/inter"
)

type JsonToJson struct{}

// Determine if json is already valid
func (j JsonToJson) IsAble(object interface{}) bool {
	var js interface{}
	result, ok := object.(string)

	return ok && json.Unmarshal([]byte(result), &js) == nil
}

func (j JsonToJson) EncodeThrough(_ inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	return object.(string), nil
}
