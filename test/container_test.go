package test

import (
	contractsHttp "github.com/lanvard/contract/http"
	"github.com/lanvard/foundation"
	http "github.com/lanvard/foundation/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testInterface interface{}
type testStruct struct {}

func Test_one_binding_with_contract(t *testing.T) {
	app := foundation.NewContainer()

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		http.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 1)
}

func Test_multiple_binding_with_contract(t *testing.T) {
	app := foundation.NewContainer()

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
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
		(*contractsHttp.Kernel)(nil),
		http.Kernel{},
	)

	app.Singleton(
		(*contractsHttp.Kernel)(nil),
		http.Kernel{},
	)

	assert.Len(t, app.GetBindings(), 1)
}

func Test_make_from_interface(t *testing.T) {
	container := foundation.NewContainer()

	container.Singleton(
		(*contractsHttp.Kernel)(nil),
		http.Kernel{},
	)

	kernel := container.Make((*contractsHttp.Kernel)(nil))

	assert.Equal(t, http.Kernel{}, kernel)
}

func Test_resolve_automatically(t *testing.T) {
	container := foundation.NewContainer()

	resolvedStruct := container.Make(testStruct{})

	assert.Equal(t, testStruct{}, resolvedStruct)
}
