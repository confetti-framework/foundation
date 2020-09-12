package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

type JsonReaderToJson struct{}

func (j JsonReaderToJson) IsAble(object interface{}) bool {
	jsonReader, ok := object.(inter.JsonReader)
	return ok && InterfaceToJson{}.IsAble(jsonReader.Json())
}

func (j JsonReaderToJson) EncodeThrough(object interface{}, encoders []inter.Encoder) (string, error) {
	jsonReader, ok := object.(inter.JsonReader)
	if !ok {
		return "", errors.New("can not encode to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	result := jsonReader.Json()

	// If the object is nil, we don't want to return an empty body
	if result == nil {
		return "null", nil
	}

	return EncodeThrough(result, encoders)
}
