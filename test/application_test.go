package test

import (
	"github.com/lanvard/contract/inter"
	httpFoundation "github.com/lanvard/foundation/http"
	"github.com/stretchr/testify/assert"
	"lanvard/app/http"
	"lanvard/bootstrap"
	"lanvard/config"
	"testing"
)

func Test_binding(t *testing.T) {
	app := bootstrap.NewApp()

	app.Container.Singleton(
		(*inter.HttpKernel)(nil),
		http.NewKernel(app),
	)

	app.Container.Singleton(
		"testSingleton",
		"testSingletonValue",
	)

	assert.GreaterOrEqual(t, len(app.Container.GetBindings()), 3)
}

func Test_application_make(t *testing.T) {
	app := bootstrap.NewApp()

	app.Container.Singleton(
		(*interfaceHttp.Kernel)(nil),
		httpFoundation.Kernel{},
	)

	kernel := app.Container.Make((*interfaceHttp.Kernel)(nil))

	assert.Equal(t, httpFoundation.Kernel{}, kernel)
}

func Test_application_environment(t *testing.T) {
	config.App.Env = "local"

	app := bootstrap.NewApp()

	environment, _ := app.Environment()
	assert.Equal(t, "local", environment)
}

func Test_application_environment_error(t *testing.T) {
	config.App.Env = ""

	app := bootstrap.NewApp()

	_, err := app.Environment()

	//noinspection GoNilness
	assert.Equal(t, "environment not found", err.Error())
}

func Test_application_is_environment(t *testing.T) {
	config.App.Env = "local"

	app := bootstrap.NewApp()

	assert.True(t, true, app.IsEnvironment("local"))
	assert.True(t, true, app.IsEnvironment("production", "local"))
	assert.False(t, false, app.IsEnvironment("production"))
	assert.False(t, false, app.IsEnvironment("production", "staging"))
}
