package lifecycle

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/support"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_one_binding_with_an_alias(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		"testStruct",
		testStruct{},
	)

	kernel := container.Make("testStruct").(testStruct)

	require.Equal(t, testStruct{}, kernel)
}

func Test_make_from_singleton(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		(*inter.HttpKernel)(nil),
		testStruct{},
	)

	kernel := container.Make((*inter.HttpKernel)(nil)).(testStruct)

	require.Equal(t, testStruct{}, kernel)
}

func Test_make_from_singleton_with_callback(t *testing.T) {
	app := foundation.NewApp()

	app.Singleton(
		testStruct{},
		func() interface{} {
			return testStruct{TestCount: 1}
		},
	)

	newStruct := app.Make(testStruct{})

	require.Equal(t, testStruct{TestCount: 1}, newStruct)
}

func Test_make_from_singleton_with_callback_specific_return_type(t *testing.T) {
	app := foundation.NewApp()

	app.Singleton(
		testStruct{},
		func() testStruct {
			return testStruct{TestCount: 1}
		},
	)

	newStruct := app.Make(testStruct{})

	require.Equal(t, testStruct{TestCount: 1}, newStruct)
}

func Test_make_from_binding_with_callback(t *testing.T) {
	app := foundation.NewApp()

	app.Bind(
		"a_callback",
		func() testStruct {
			return testStruct{TestCount: 1}
		},
	)

	newStruct := app.Make("a_callback")

	require.Equal(t, reflect.Func, support.Kind(newStruct))
	require.Equal(t, 1, newStruct.(func() testStruct)().TestCount)
}

func Test_make_from_singleton_with_parameters(t *testing.T) {
	app := foundation.NewApp()

	app.Singleton(
		"a_callback",
		func(firstPar string) testStruct {
			return testStruct{TestCount: 1}
		},
	)

	newStruct, err := app.MakeE("a_callback")

	require.NotNil(t, err)
	require.Nil(t, newStruct)
	require.EqualError(t, err, "get instance 'a_callback' from container: Can not instantiate callback with parameters")
}

func Test_resolve_automatically(t *testing.T) {
	container := foundation.NewContainer()

	resolvedStruct := container.Make(testStruct{})

	require.Equal(t, testStruct{}, resolvedStruct)
}

func Test_binding_existing_object(t *testing.T) {
	container := foundation.NewContainer()

	kernel := testStruct{}
	container.Bind("testStruct", kernel)

	resolvedStruct := container.Make("testStruct")

	require.Equal(t, testStruct{}, resolvedStruct)
}

func Test_one_binding_with_contract(t *testing.T) {
	container := foundation.NewContainer()

	container.Bind(
		(*inter.HttpKernel)(nil),
		testStruct{},
	)

	require.Len(t, container.Bindings(), 1)
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

	require.Len(t, app.Bindings(), 2)
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

	require.Len(t, container.Bindings(), 1)
}

func Test_binding_and_make_from_interface(t *testing.T) {
	container := foundation.NewContainer()

	kernel := testStruct{}
	container.Bind((*testInterface)(nil), kernel)

	resolvedStruct := container.Make((*testInterface)(nil)).(testInterface)

	require.Equal(t, testStruct{}, resolvedStruct)
}

func Test_binding_without_abstract(t *testing.T) {
	container := foundation.NewContainer()

	container.Instance(testStruct{TestCount: 1})

	resolvedStruct := container.Make(testStruct{}).(testStruct)

	require.Equal(t, testStruct{TestCount: 1}, resolvedStruct)
}

func Test_binding_missing_without_abstract(t *testing.T) {
	container := foundation.NewContainer()

	resolvedStruct := container.Make(testStruct{}).(testStruct)

	require.Equal(t, testStruct{TestCount: 0}, resolvedStruct)
}

func Test_resolve_with_an_pointer_value_and_empty_instance(t *testing.T) {
	container := foundation.NewContainer()

	instance := testStruct{}
	container.Instance(instance)

	var requestedStruct testStruct
	_, err := container.MakeE(&requestedStruct)

	require.Nil(t, err)
	require.Equal(t, testStruct{TestCount: 0}, requestedStruct)
}

func Test_resolve_with_an_pointer_value(t *testing.T) {
	container := foundation.NewContainer()

	instance := testStruct{TestCount: 1}
	container.Instance(instance)

	var requestedStruct testStruct
	_, err := container.MakeE(&requestedStruct)

	require.Nil(t, err)
	require.Equal(t, testStruct{TestCount: 1}, requestedStruct)
}

func Test_extending_bindings(t *testing.T) {

	container := foundation.NewContainer()

	container.Bind(testStruct{}, testStruct{TestCount: 1})

	container.Extend(testStruct{}, func(service interface{}) interface{} {
		testService := service.(testStruct)
		newVariable := testService.TestCount + 2

		return testStruct{TestCount: newVariable}
	})

	resolvedStruct := container.Make(testStruct{}).(testStruct)

	require.Equal(t, testStruct{TestCount: 3}, resolvedStruct)
}

func Test_resolve_with_boot_app(t *testing.T) {
	bootContainer := foundation.NewContainer()
	container := foundation.NewContainerByBoot(bootContainer)

	container.Bind("application_name", "Cooler")

	require.Equal(t, "Cooler", container.Make("application_name"))
}

func Test_resolve_from_boot_app(t *testing.T) {
	bootContainer := foundation.NewContainer()
	bootContainer.Bind("application_name", "Cooler")

	container := foundation.NewContainerByBoot(bootContainer)

	require.Equal(t, "Cooler", container.Make("application_name"))
}

func Test_resolve_from_normal_container_but_not_from_boot(t *testing.T) {
	bootContainer := foundation.NewContainer()
	bootContainer.Bind("application_name", "Heater")

	container := foundation.NewContainerByBoot(bootContainer)
	container.Bind("application_name", "Cooler")

	require.Equal(t, "Cooler", container.Make("application_name"))
}
