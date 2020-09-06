package transformer

import "github.com/lanvard/contract/inter"

type JsonReaderToJson struct{}

func (j JsonReaderToJson) Transformable(object interface{}) bool {
	_, ok := object.(inter.JsonReader)
	return ok
}

func (j JsonReaderToJson) Transform(object interface{}) string {
	_, ok := object.(inter.JsonReader)
	if !ok {
		panic("can not transform to json with an unsupported type")
	}

	panic("test")
}
