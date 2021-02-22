package service

import (
	"reflect"
)

type ParsedOption struct {
	Number int
	Tag    reflect.StructTag
	Value  interface{}
}

func GetOptions(command interface{}) []ParsedOption {
	//goland:noinspection GoPreferNilSlice
	result := []ParsedOption{}
	elem := reflect.ValueOf(command)

	for i := 0; i < elem.NumField(); i++ {
		typeField := reflect.TypeOf(command).Field(i)
		tag := typeField.Tag
		field := elem.Field(i)
		option := ParsedOption{
			Number: i,
			Tag:    tag,
			Value:  field.Interface(),
		}
		result = append(result, option)
	}

	return result
}
