package transformer

import (
	"github.com/lanvard/contract/inter"
	"reflect"
)

type JsonReaderToJson struct{}

func (j JsonReaderToJson) Transformable(object interface{}) bool {
	jsonReader, ok := object.(inter.JsonReader)
	return ok && InterfaceToJson{}.Transformable(jsonReader.Json())
}

func (j JsonReaderToJson) Transform(object interface{}) string {
	jsonReader, ok := object.(inter.JsonReader)
	if !ok {
		panic("can not transform to json with an unsupported type " + reflect.TypeOf(object).String())
	}

	if (ValueToJson{}.Transformable(jsonReader.Json())) {
		return ValueToJson{}.Transform(jsonReader.Json())
	}

	return InterfaceToJson{}.Transform(jsonReader.Json())
}
