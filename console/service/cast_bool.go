package service

import "github.com/confetti-framework/support"

func CastToBool(input string) interface{} {
	return support.NewValue(input).Bool()
}
