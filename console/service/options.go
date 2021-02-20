package service

import (
	"github.com/confetti-framework/support"
	"reflect"
)

type ParsedOption struct {
	Number   int
	Tag      reflect.StructTag
	TypeName string
}

type ValueOption struct {
	Value  bool
	Option ParsedOption
}

func GetOptions(command interface{}) []ParsedOption {
	//goland:noinspection GoPreferNilSlice
	result := []ParsedOption{}
	elem := reflect.ValueOf(command)

	for i := 0; i < elem.NumField(); i++ {
		typeField := reflect.TypeOf(command).Field(i)
		tag := typeField.Tag
		field := elem.Field(i)
		result = append(result, ParsedOption{Number: i, Tag: tag, TypeName: support.Name(field.Interface())})
	}

	return result
}
