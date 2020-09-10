package transformer

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
	"reflect"
)

type ValueToJson struct{}

func (v ValueToJson) Transformable(object interface{}) bool {
	value, ok := object.(support.Value)

	return ok && InterfaceToJson{}.Transformable(value.Raw())
}

func (v ValueToJson) TransformThrough(object interface{}, encoders []inter.ResponseEncoder) (string, error) {
	value, ok := object.(support.Value)
	if !ok {
		return "", errors.New("can not transform to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	return TransformThrough(value.Raw(), encoders)
}
