package encoder

import (
	"github.com/confetti-framework/contract/inter"
	"reflect"
)

type StringToString struct{}

func (j StringToString) IsAble(object interface{}) bool {
	_, ok := object.(string)
	return ok || object == nil
}

func (j StringToString) EncodeThrough(_ inter.App, object interface{}, _ []inter.Encoder) (string, error) {
	if object == nil {
		return "", nil
	}

	result, ok := object.(string)
	if !ok {
		return "", EncodeError.Wrap("can not encode to string with an unsupported type " + reflect.TypeOf(object).String())
	}

	return result, nil
}
