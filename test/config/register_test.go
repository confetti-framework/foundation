package config

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/providers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapOfNilIndex(t *testing.T) {
	var index []interface{}
	result := providers.ConfigServiceProvider{Index: index}.GetMap()

	assert.Equal(t, map[string]interface{}{}, result)
}

func TestMapOfStruct(t *testing.T) {
	config := firstConfig{Title: "OK"}
	var index = []interface{}{config}

	result := providers.ConfigServiceProvider{Index: index}.GetMap()

	assert.Equal(t, map[string]interface{}{"firstConfig": config}, result)
}

func TestRegisterNilIndex(t *testing.T) {
	var index []interface{}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(t, map[string]interface{}{}, container.Make("config"))
}

func TestRegisterStruct(t *testing.T) {
	index := []interface{}{firstConfig{Title: "Sock"}}
	var container inter.Container = foundation.NewContainer()
	container = providers.ConfigServiceProvider{Index: index}.Register(container)

	assert.Equal(
		t,
		firstConfig{Title: "Sock"},
		container.Make("config.firstConfig"),
	)
}

func TestSetAnonymousStruct(t *testing.T) {
	var config = struct{ title string }{title: "The horse"}
	index := []interface{}{config}
	var container inter.Container = foundation.NewContainer()
	assert.PanicsWithError(t, "config need to be a named struct. struct { title string } given.", func() {
		providers.ConfigServiceProvider{Index: index}.Register(container)
	})
}

type firstConfig struct {
	Title string
}

type secondConfig struct{}

type deepConfig struct {
	Deep map[string]interface{}
}
