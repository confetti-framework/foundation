package config

import (
	"github.com/lanvard/foundation/providers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterNil(t *testing.T) {
	var index []interface{}

	result := providers.ConfigServiceProvider{Index: index}.GetMap()

	assert.Equal(t, map[string]interface{}{}, result)
}

func TestRegisterStruct(t *testing.T) {
	config := FirstConfig{Title: "OK"}
	var index = []interface{}{config}

	result := providers.ConfigServiceProvider{Index: index}.GetMap()

	assert.Equal(t, map[string]interface{}{"FirstConfig": config}, result)
}

type FirstConfig struct {
	Title string
}
