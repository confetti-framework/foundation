package transformer

import (
	"github.com/lanvard/support"
	"reflect"
)

type ValueToJson struct{}

func (v ValueToJson) Transformable(object interface{}) bool {
	value, ok := object.(support.Value)

	return ok && InterfaceToJson{}.Transformable(value.Raw())
}

func (v ValueToJson) Transform(object interface{}) (string, error) {
	value, ok := object.(support.Value)
	if !ok {
		panic("can not transform to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	return InterfaceToJson{}.Transform(value.Raw())
}
