package config

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/providers"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_get_empty_struct_by_empty_key(t *testing.T) {
	var index = map[string]interface{}{}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	result := container.Make("config.")
	require.Equal(t, map[string]interface{}{}, result)
}

func Test_get_filled_struct_by_empty_key(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	require.Equal(t, map[string]interface{}{"firstConfig": firstConfig}, container.Make("config"))
}

func Test_get_by_not_existing_key(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	result, err := container.MakeE("config.age")
	require.Nil(t, result)
	require.EqualError(t, err, "get instance 'config.age' from container: key 'age': can not found value")
}

func Test_get_by_existing_key(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	require.Equal(t, "The horse", container.Make("config.firstConfig.Title"))
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

	require.Equal(t, "bottom of the sea", container.Make("config.deepConfig.Deep.deeper.0"))
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

	require.Equal(t, "bottom of the sea", container.Make("config.deepConfig.Deep.deeper.0"))
}

func Test_get_by_asterisk(t *testing.T) {
	var index = map[string]interface{}{"firstConfig": firstConfig, "secondConfig": secondConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	require.Len(t, container.Make("config.*"), 2)
}
