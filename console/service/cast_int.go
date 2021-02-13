package service

import "github.com/spf13/cast"

func CastToInt(input string) interface{} {
	return cast.ToInt(input)
}
