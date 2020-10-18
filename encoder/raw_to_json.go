package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"reflect"
)

type RawToJson struct{}

func (v RawToJson) IsAble(object interface{}) bool {
	value, ok := object.(interface{ Raw() interface{} })
	return ok && InterfaceToJson{}.IsAble(value.Raw())
}

func (v RawToJson) EncodeThrough(app inter.App, object interface{}, encoders []inter.Encoder) (string, error) {
	value, ok := object.(interface{ Raw() interface{} })
	if !ok {
		return "", errors.New("can not encode to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	result := value.Raw()

	// If the object is nil, we don't want to return an empty body
	if result == nil {
		return "null", nil
	}

	return EncodeThrough(app, result, encoders)
}
