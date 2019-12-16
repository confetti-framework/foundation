package test

import (
	interfaceHttp "github.com/lanvard/contract/http"
	httpFoundation "github.com/lanvard/foundation/http"
	"github.com/stretchr/testify/assert"
	"lanvard/app/http"
	"lanvard/bootstrap"
	"testing"
)

func Test_binding(t *testing.T) {
	app := bootstrap.NewApp()

	app.Container.Singleton(
		(*interfaceHttp.Kernel)(nil),
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
