package test

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/http"
	"github.com/stretchr/testify/assert"
	"lanvard/bootstrap"
	"testing"
)

type testInterface interface{}
type testStruct struct {
	App *foundation.Application
	TestCount int
}

func Test_one_binding_with_an_alias(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		"http.Kernel",
		http.Kernel{},
	)

	kernel := container.Make("http.Kernel").(http.Kernel)

	assert.Equal(t, http.Kernel{}, kernel)
}

func Test_make_from_singleton(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		(*inter.HttpKernel)(nil),
		http.Kernel{},
	)

	kernel := container.Make((*inter.HttpKernel)(nil)).(http.Kernel)

	assert.Equal(t, http.Kernel{}, kernel)
}

func Test_make_from_singleton_with_callback(t *testing.T) {
	app := foundation.Application{
		Container: foundation.NewContainer(),
	}

	app.Container().Singleton(
		testStruct{},
		func() interface{} {
			return testStruct{TestCount: 1}
		},
	)

	newStruct := app.Make(testStruct{}).(testStruct)

	assert.Equal(t, testStruct{TestCount: 1}, newStruct)
}

func Test_resolve_automatically(t *testing.T) {
	container := foundation.NewContainer()

	resolvedStruct := container.Make(testStruct{})

	assert.Equal(t, testStruct{}, resolvedStruct)
}

func Test_binding_existing_object(t *testing.T) {
	container := foundation.NewContainer()

	kernel := testStruct{}
	container.Instance("http.Kernel", kernel)

	resolvedStruct := container.Make("http.Kernel")

	assert.Equal(t, testStruct{}, resolvedStruct)
}

func Test_one_binding_with_contract(t *testing.T) {
	app := foundation.NewContainer()

	app.Bind(
		(*inter.HttpKernel)(nil),
		http.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 1)
}

func Test_multiple_binding_with_contract(t *testing.T) {
	app := foundation.NewContainer()

	app.Bind(
		(*inter.HttpKernel)(nil),
		http.Kernel{},
	)

	app.Bind(
		(*testInterface)(nil),
		http.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 2)
}

func Test_binding_two_with_the_same_interfaces(t *testing.T) {
	app := foundation.NewContainer()

	app.Bind(
		(*inter.HttpKernel)(nil),
		http.Kernel{},
	)

	app.Bind(
		(*inter.HttpKernel)(nil),
		http.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 1)
}

func Test_binding_and_make_from_interface(t *testing.T) {
	container := foundation.NewContainer()

	kernel := testStruct{}
	container.Instance((*testInterface)(nil), kernel)

	resolvedStruct := container.Make((*testInterface)(nil)).(testInterface)

	assert.Equal(t, testStruct{}, resolvedStruct)
}

func Test_binding_without_abstract(t *testing.T) {
	container := foundation.NewContainer()

	container.JustBind(testStruct{TestCount: 1})

	resolvedStruct := container.Make(testStruct{}).(testStruct)

	assert.Equal(t, testStruct{TestCount: 1}, resolvedStruct)
}

func Test_extending_bindings(t *testing.T) {

	app := bootstrap.NewApp()

	app.Container.Instance(testStruct{}, testStruct{TestCount: 1})

	app.Container.Extend(testStruct{}, func(service interface{}) interface{} {
		testService := service.(testStruct)
		newVariable := testService.TestCount + 2

		return testStruct{App: app, TestCount: newVariable}
	})

	resolvedStruct := app.Make(testStruct{}).(testStruct)

	assert.Equal(t, testStruct{App: app, TestCount: 3}, resolvedStruct)
}
