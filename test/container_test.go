package test

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOneBindingWithAnAlias(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		"testStruct",
		testStruct{},
	)

	kernel := container.Make("testStruct").(testStruct)

	assert.Equal(t, testStruct{}, kernel)
}

func TestMakeFromSingleton(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		(*inter.HttpKernel)(nil),
		testStruct{},
	)

	kernel := container.Make((*inter.HttpKernel)(nil)).(testStruct)

	assert.Equal(t, testStruct{}, kernel)
}

func TestMakeFromSingletonWithCallback(t *testing.T) {
	app := foundation.NewApp()

	app.Singleton(
		testStruct{},
		func() interface{} {
			return testStruct{TestCount: 1}
		},
	)

	newStruct := app.Make(testStruct{})

	assert.Equal(t, testStruct{TestCount: 1}, newStruct)
}

func TestResolveAutomatically(t *testing.T) {
	container := foundation.NewContainer()

	resolvedStruct := container.Make(testStruct{})

	assert.Equal(t, testStruct{}, resolvedStruct)
}

func TestBindingExistingObject(t *testing.T) {
	container := foundation.NewContainer()

	kernel := testStruct{}
	container.Instance("testStruct", kernel)

	resolvedStruct := container.Make("testStruct")

	assert.Equal(t, testStruct{}, resolvedStruct)
}

func TestOneBindingWithContract(t *testing.T) {
	container := foundation.NewContainer()

	container.Bind(
		(*inter.HttpKernel)(nil),
		testStruct{},
	)

	assert.Len(t, container.Bindings(), 1)
}

func TestMultipleBindingWithContract(t *testing.T) {
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

func TestBindingTwoWithTheSameInterfaces(t *testing.T) {
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

func TestBindingAndMakeFromInterface(t *testing.T) {
	container := foundation.NewContainer()

	kernel := testStruct{}
	container.Instance((*testInterface)(nil), kernel)

	resolvedStruct := container.Make((*testInterface)(nil)).(testInterface)

	assert.Equal(t, testStruct{}, resolvedStruct)
}

func TestBindingWithoutAbstract(t *testing.T) {
	container := foundation.NewContainer()

	container.BindStruct(testStruct{TestCount: 1})

	resolvedStruct := container.Make(testStruct{}).(testStruct)

	assert.Equal(t, testStruct{TestCount: 1}, resolvedStruct)
}

func TestExtendingBindings(t *testing.T) {

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

func TestResolveWithBootApp(t *testing.T) {
	bootContainer := foundation.NewContainer()
	container := foundation.NewContainerByBoot(bootContainer)

	container.Instance("application_name", "Cooler")

	assert.Equal(t, "Cooler", container.Make("application_name"))
}

func TestResolveFromBootApp(t *testing.T) {
	bootContainer := foundation.NewContainer()
	bootContainer.Instance("application_name", "Cooler")

	container := foundation.NewContainerByBoot(bootContainer)

	assert.Equal(t, "Cooler", container.Make("application_name"))
}

func TestResolveFromNormalContainerButNotFromBoot(t *testing.T) {
	bootContainer := foundation.NewContainer()
	bootContainer.Instance("application_name", "Heater")

	container := foundation.NewContainerByBoot(bootContainer)
	container.Instance("application_name", "Cooler")

	assert.Equal(t, "Cooler", container.Make("application_name"))
}
