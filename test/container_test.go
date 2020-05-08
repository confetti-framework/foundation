package test

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_one_binding_with_an_alias(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		"testStruct",
		testStruct{},
	)

	kernel := container.Make("testStruct").(testStruct)

	assert.Equal(t, testStruct{}, kernel)
}

func Test_make_from_singleton(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		(*inter.HttpKernel)(nil),
		testStruct{},
	)

	kernel := container.Make((*inter.HttpKernel)(nil)).(testStruct)

	assert.Equal(t, testStruct{}, kernel)
}

func Test_make_from_singleton_with_callback(t *testing.T) {
	app := foundation.NewApp()

	app.Singleton(
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
	container.Instance("testStruct", kernel)

	resolvedStruct := container.Make("testStruct")

	assert.Equal(t, testStruct{}, resolvedStruct)
}

func Test_one_binding_with_contract(t *testing.T) {
	container := foundation.NewContainer()

	container.Bind(
		(*inter.HttpKernel)(nil),
		testStruct{},
	)

	assert.Len(t, container.Bindings(), 1)
}

func Test_multiple_binding_with_contract(t *testing.T) {
	app := foundation.NewContainer()

	app.Bind(
		(*inter.HttpKernel)(nil),
		testStruct{},
	)

	app.Bind(
		(*testInterface)(nil),
		testStruct{},
	)

	assert.Len(t, app.Bindings(), 2)
}

func Test_binding_two_with_the_same_interfaces(t *testing.T) {
	container := foundation.NewContainer()

	container.Bind(
		(*inter.HttpKernel)(nil),
		testStruct{},
	)

	container.Bind(
		(*inter.HttpKernel)(nil),
		testStruct{},
	)

	assert.Len(t, container.Bindings(), 1)
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

	container.BindStruct(testStruct{TestCount: 1})

	resolvedStruct := container.Make(testStruct{}).(testStruct)

	assert.Equal(t, testStruct{TestCount: 1}, resolvedStruct)
}

func Test_extending_bindings(t *testing.T) {

	container := foundation.NewContainer()

	container.Instance(testStruct{}, testStruct{TestCount: 1})

	container.Extend(testStruct{}, func(service interface{}) interface{} {
		testService := service.(testStruct)
		newVariable := testService.TestCount + 2

		return testStruct{TestCount: newVariable}
	})

	resolvedStruct := container.Make(testStruct{}).(testStruct)

	assert.Equal(t, testStruct{TestCount: 3}, resolvedStruct)
}
