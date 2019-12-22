package test

import (
	contract "github.com/lanvard/contract/http"
	"github.com/lanvard/foundation"
	http "github.com/lanvard/foundation/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testInterface interface{}
type testStruct struct {
	TestCount int
}

func Test_one_binding_with_contract(t *testing.T) {
	app := foundation.NewContainer()

	app.Singleton(
		(*contract.Kernel)(nil),
		http.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 1)
}

func Test_multiple_binding_with_contract(t *testing.T) {
	app := foundation.NewContainer()

	app.Singleton(
		(*contract.Kernel)(nil),
		http.Kernel{},
	)

	app.Singleton(
		(*testInterface)(nil),
		http.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 2)
}

func Test_binding_two_with_the_same_interfaces(t *testing.T) {
	app := foundation.NewContainer()

	app.Singleton(
		(*contract.Kernel)(nil),
		http.Kernel{},
	)

	app.Singleton(
		(*contract.Kernel)(nil),
		http.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 1)
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
		(*contract.Kernel)(nil),
		http.Kernel{},
	)

	kernel := container.Make((*contract.Kernel)(nil)).(http.Kernel)

	assert.Equal(t, http.Kernel{}, kernel)
}

func Test_make_from_singleton_with_callback(t *testing.T) {
	app := foundation.Application{
		Container: foundation.NewContainer(),
	}

	app.Container.Singleton(
		testStruct{},
		func() interface{} {
			return testStruct{}
		},
	)

	newStruct := app.Make(testStruct{}).(testStruct)

	assert.Equal(t, testStruct{}, newStruct)
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
