package config

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/providers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_get_empty_struct_by_empty_key(t *testing.T) {
	var index = map[string]interface{}{}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(
		t,
		map[string]interface{}{},
		container.Make("config."),
	)
}

func Test_get_filled_struct_by_empty_key(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(t, map[string]interface{}{"firstConfig": firstConfig}, container.Make("config"))
}

func Test_get_by_not_existing_key(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	result, err := container.MakeE("config.age")
	assert.Nil(t, result)
	assert.EqualError(t, err, "no value found with key 'age'")
}

func Test_get_by_existing_key(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(t, "The horse", container.Make("config.firstConfig.Title"))
}

func Test_get_nested_from_struct(t *testing.T) {
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

func Test_get_from_booted_container(t *testing.T) {
	index := map[string]interface{}{"deepConfig": deepConfig{
		Deep: map[string]interface{}{
			"deeper": []string{
				"bottom of the sea",
			},
		},
	}}
	var bootContainer inter.Container = foundation.NewContainer()
	bootContainer = providers.ConfigServiceProvider{Index: index}.Register(bootContainer)
	container := foundation.NewContainerByBoot(bootContainer)

	assert.Equal(t, "bottom of the sea", container.Make("config.deepConfig.Deep.deeper.0"))
}

func Test_get_by_asterisk(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig, "secondConfig": secondConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Len(t, container.Make("config.*"), 2)
}
