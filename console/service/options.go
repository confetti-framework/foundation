package service

import (
	"github.com/confetti-framework/contract/inter"
	"reflect"
)

type ParsedOption struct {
	Number int
	Tag    reflect.StructTag
	Caster inter.Caster
}

func GetOptions(command interface{}) []ParsedOption {
	//goland:noinspection GoPreferNilSlice
	result := []ParsedOption{}
	elem := reflect.ValueOf(command)

	for i := 0; i < elem.NumField(); i++ {
		tag := reflect.TypeOf(command).Field(i).Tag
		field := elem.Field(i)
		caster := getCasterByKind(field.Kind())
		result = append(result, ParsedOption{Number: i, Tag: tag, Caster: caster})
	}

	return result
}

func getCasterByKind(kind reflect.Kind) inter.Caster {
	switch kind {
	case reflect.Bool:
		return CastToBool
	case reflect.String:
		return CastToString
	case reflect.Int:
		return CastToInt
	case reflect.Float64:
		return CastToFloat
	}
	return nil
}
