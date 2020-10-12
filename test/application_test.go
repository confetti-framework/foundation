package test

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testInterface interface{}
type testStruct struct {
	App       *inter.App
	TestCount int
}

func Test_binding(t *testing.T) {
	var container inter.Container = foundation.NewContainer()

	app := &foundation.Application{}
	app.SetContainer(container)

	app.Singleton(
		(*testInterface)(nil),
		testStruct{},
	)

	app.Singleton(
		"testSingleton",
		"testSingletonValue",
	)

	assert.Equal(t, len((*app.Container()).Bindings()), 2)
}

func Test_application_make(t *testing.T) {
	app := foundation.NewApp()

	(*app.Container()).Singleton(
		(*testInterface)(nil),
		testStruct{},
	)

	struct1 := app.Make((*testInterface)(nil))

	assert.Equal(t, testStruct{}, struct1)
}

func Test_application_environment(t *testing.T) {
	app := foundation.NewApp()
	(*app.Container()).Bind("env", "local")

	environment, _ := app.Environment()
	assert.Equal(t, "local", environment)
}

func Test_application_environment_error(t *testing.T) {
	app := foundation.NewApp()
	(*app.Container()).Bind("env", "")

	_, err := app.Environment()

	//noinspection GoNilness
	assert.Equal(t, "environment not found", err.Error())
}

func Test_application_is_environment(t *testing.T) {
	app := foundation.NewApp()
	(*app.Container()).Bind("env", "local")

	assert.True(t, true, app.IsEnvironment("local"))
	assert.True(t, true, app.IsEnvironment("production", "local"))
	assert.False(t, false, app.IsEnvironment("production"))
	assert.False(t, false, app.IsEnvironment("production", "acceptance"))
}
