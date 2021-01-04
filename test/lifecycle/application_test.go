package lifecycle

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/stretchr/testify/require"
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

	app.Bind(
		(*testInterface)(nil),
		testStruct{},
	)

	app.Bind(
		"testSingleton",
		"testSingletonValue",
	)

	require.Equal(t, len((*app.Container()).Bindings()), 2)
}

func Test_binding_from_boot_container_and_application(t *testing.T) {
	var bootContainer inter.Container = foundation.NewContainer()

	bootContainer.Bind(
		(*testInterface)(nil),
		testStruct{},
	)
	container := foundation.NewContainerByBoot(bootContainer)

	container.Bind(
		"testSingleton",
		"testSingletonValue",
	)

	app := &foundation.Application{}
	app.SetContainer(container)

	require.Equal(t, len((*app.Container()).Bindings()), 2)
}

func Test_singleton_from_container(t *testing.T) {
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

	require.Equal(t, len((*app.Container()).Bindings()), 2)
}

func Test_singleton_with_simple_string_callback(t *testing.T) {
	var bootContainer inter.Container = foundation.NewContainer()

	bootContainer.Singleton(
		"current_user",
		func() string {
			return "Joop"
		},
	)

	appFromJoop := &foundation.Application{}
	appFromJoop.SetContainer(foundation.NewContainerByBoot(bootContainer))

	require.Equal(t, "Joop", appFromJoop.Make("current_user"))
}

func Test_singleton_with_multiple_requests(t *testing.T) {
	var bootContainer inter.Container = foundation.NewContainer()
	var currentUser = "None"

	bootContainer.Singleton(
		"current_user",
		func() string {
			return currentUser
		},
	)

	// First request
	var containerFromJoop = foundation.NewContainerByBoot(bootContainer)
	appFromJoop := &foundation.Application{}
	appFromJoop.SetContainer(containerFromJoop)
	currentUser = "Joop"
	expectJoop := appFromJoop.Make("current_user")

	// Second request
	var containerFromPiet = foundation.NewContainerByBoot(bootContainer)
	appFromPiet := &foundation.Application{}
	appFromPiet.SetContainer(containerFromPiet)
	currentUser = "Piet"
	expectPiet := containerFromPiet.Make("current_user")

	require.Equal(t, "Joop", expectJoop)
	require.Equal(t, "Piet", expectPiet)
}

func Test_singleton_in_boot_container_within_same_requests(t *testing.T) {
	var bootContainer inter.Container = foundation.NewContainer()
	var currentUser = "None"

	bootContainer.Singleton(
		"current_user",
		func() string {
			return currentUser
		},
	)

	var containerFromJoop = foundation.NewContainerByBoot(bootContainer)
	appFromJoop := &foundation.Application{}
	appFromJoop.SetContainer(containerFromJoop)
	currentUser = "Joop"
	expectJoop := appFromJoop.Make("current_user")
	currentUser = "Piet"
	expectJoop = appFromJoop.Make("current_user")

	require.Equal(t, "Joop", expectJoop)
}

func Test_singleton_within_same_requests(t *testing.T) {
	var bootContainer inter.Container = foundation.NewContainer()
	var currentUser = "None"

	var containerFromJoop = foundation.NewContainerByBoot(bootContainer)

	containerFromJoop.Singleton(
		"current_user",
		func() string {
			return currentUser
		},
	)
	appFromJoop := &foundation.Application{}
	appFromJoop.SetContainer(containerFromJoop)
	currentUser = "Joop"
	expectJoop := appFromJoop.Make("current_user")
	currentUser = "Piet"
	expectJoop = appFromJoop.Make("current_user")

	require.Equal(t, "Joop", expectJoop)
}

func Test_application_make(t *testing.T) {
	app := foundation.NewApp()

	(*app.Container()).Singleton(
		(*testInterface)(nil),
		testStruct{},
	)

	struct1 := app.Make((*testInterface)(nil))

	require.Equal(t, testStruct{}, struct1)
}

func Test_application_environment(t *testing.T) {
	app := foundation.NewApp()
	(*app.Container()).Bind("env", "local")

	environment, _ := app.Environment()
	require.Equal(t, "local", environment)
}

func Test_application_environment_error(t *testing.T) {
	app := foundation.NewApp()
	(*app.Container()).Bind("env", "")

	_, err := app.Environment()

	//noinspection GoNilness
	require.Equal(t, "environment not found", err.Error())
}

func Test_application_is_environment(t *testing.T) {
	app := foundation.NewApp()
	(*app.Container()).Bind("env", "local")

	require.True(t, true, app.IsEnvironment("local"))
	require.True(t, true, app.IsEnvironment("production", "local"))
	require.False(t, false, app.IsEnvironment("production"))
	require.False(t, false, app.IsEnvironment("production", "acceptance"))
}
