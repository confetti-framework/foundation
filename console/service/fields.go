package service

import (
	"reflect"
)

type Field struct {
	Number int
	Tag    reflect.StructTag
	Value  interface{}
}

func GetCommandFields(command interface{}) []Field {
	//goland:noinspection GoPreferNilSlice
	result := []Field{}
	elem := reflect.ValueOf(command)

	for i := 0; i < elem.NumField(); i++ {
		typeField := reflect.TypeOf(command).Field(i)
		tag := typeField.Tag
		field := elem.Field(i)
		result = append(result, Field{
			Number: i,
			Tag:    tag,
			Value:  field.Interface(),
		})
	}

	return result
}
