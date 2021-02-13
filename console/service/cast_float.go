package service

import "github.com/spf13/cast"

func CastToFloat(input string) interface{} {
	return cast.ToFloat64(input)
}
