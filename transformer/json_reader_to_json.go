package transformer

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

type JsonReaderToJson struct{}

func (j JsonReaderToJson) Transformable(object interface{}) bool {
	jsonReader, ok := object.(inter.JsonReader)
	return ok && InterfaceToJson{}.Transformable(jsonReader.Json())
}

func (j JsonReaderToJson) TransformThrough(object interface{}, encoders []inter.ResponseEncoder) (string, error) {
	jsonReader, ok := object.(inter.JsonReader)
	if !ok {
		return "", errors.New("can not transform to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	return TransformThrough(jsonReader.Json(), encoders)
}
