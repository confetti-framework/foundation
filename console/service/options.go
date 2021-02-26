package service

import (
	"reflect"
)

type Field struct {
	Number int
	Tag    reflect.StructTag
	Value  interface{}
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
			Value:  field.Interface(),
		}
		result = append(result, option)
	}

	return result
}
