package config

import (
	"github.com/lanvard/support"
)

func Get(config interface{}, key string) interface{} {
	result, err := GetE(config, key)
	if err != nil {
		panic(err)
	}

	return result
}

func GetE(config interface{}, key string) (interface{}, error) {
	return support.NewValue(config).Get(key).RawE()
}
