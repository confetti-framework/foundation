package config

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/providers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_register_nil_index(t *testing.T) {
	var index map[string]interface{}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(t, map[string]interface{}(nil), container.Make("config"))
}

func Test_register_struct(t *testing.T) {
	index := map[string]interface{}{"firstConfig": firstConfig}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(
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
