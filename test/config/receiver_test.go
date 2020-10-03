package config

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/providers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEmptyStructByEmptyKey(t *testing.T) {
	var index = map[string]interface{}{}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(
		t,
		map[string]interface{}{},
		container.Make("config."),
	)
}

func TestGetFilledStructByEmptyKey(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(t, map[string]interface{}{"firstConfig": firstConfig}, container.Make("config"))
}

func TestGetByNotExistingKey(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(t, nil, container.Make("config.age"))
}

func TestGetByExistingKey(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(t, nil, container.Make("config.Title"))
}

func TestGetNestedFromStruct(t *testing.T) {
	index := map[string]interface{}{"deepConfig": deepConfig{
		Deep: map[string]interface{}{
			"deeper": []string{
				"bottom of the sea",
			},
		},
	}}

	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(t, "bottom of the sea", container.Make("config.deepConfig.Deep.deeper.0"))
}

func TestGetByAsterisk(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig, "secondConfig": secondConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Len(t, container.Make("config.*"), 2)
}
