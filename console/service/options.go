package service

import (
	"github.com/confetti-framework/support"
	"reflect"
)

type Field struct {
	Number int
	Tag    reflect.StructTag
	Type   string
}

func GetOptions(command interface{}) []Field {
	//goland:noinspection GoPreferNilSlice
	result := []Field{}
	elem := reflect.ValueOf(command)

	for i := 0; i < elem.NumField(); i++ {
		typeField := reflect.TypeOf(command).Field(i)
		tag := typeField.Tag
		field := elem.Field(i)
		option := Field{
			Number: i,
			Tag:    tag,
			Type:   support.Name(field.Interface()),
		}
		result = append(result, option)
	}

	return result
}
