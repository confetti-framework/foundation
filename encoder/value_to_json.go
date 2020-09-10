package encoder

import (
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support"
	"reflect"
)

type ValueToJson struct{}

func (v ValueToJson) IsAble(object interface{}) bool {
	value, ok := object.(support.Value)

	return ok && InterfaceToJson{}.IsAble(value.Raw())
}

func (v ValueToJson) EncodeThrough(object interface{}, encoders []inter.Encoder) (string, error) {
	value, ok := object.(support.Value)
	if !ok {
		return "", errors.New("can not transform to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	return TransformThrough(value.Raw(), encoders)
}
