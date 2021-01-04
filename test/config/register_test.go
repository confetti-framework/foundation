package config

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/providers"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_register_nil_index(t *testing.T) {
	var index map[string]interface{}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	require.Equal(t, map[string]interface{}(nil), container.Make("config"))
}

func Test_register_struct(t *testing.T) {
	index := map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	require.Equal(
		t,
		firstConfig,
		container.Make("config.firstConfig"),
	)
}

var firstConfig = config{Title: "The horse"}
var secondConfig = config{Title: "The Dog"}

type config struct {
	Title string
}

type deepConfig struct {
	Deep map[string]interface{}
}
