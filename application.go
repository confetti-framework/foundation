package foundation

import (
	"fmt"
	"github.com/lanvard/contract/inter"
)

type Application struct {
	// The service container
	container *inter.Container
}

func NewApp() *Application {
	app := Application{}

	container := NewContainer()
	app.SetContainer(container)

	return &app
}

// Get the service container
func (a *Application) Container() *inter.Container {
	return a.container
}

// Set the service container
func (a *Application) SetContainer(container inter.Container) {
	a.container = &container
}

// Register a shared binding in the container.
func (a *Application) Singleton(abstract interface{}, concrete interface{}) {
	(*a.Container()).Singleton(abstract, concrete)
}

// Resolve the given type from the container.
func (a *Application) Make(abstract interface{}) interface{} {
	return (*a.container).Make(abstract)
}

func (a *Application) Instance(abstract interface{}, concrete interface{}) interface{} {
	return (*a.container).Instance(abstract, concrete)
}

// Bind all of the application paths in the container.
func (a *Application) BindPathsInContainer(path inter.BasePath) {
	container := *a.container
	container.Instance("path.app", path.AppPath())
	container.Instance("path.base", path.BasePath())
	container.Instance("path.lang", path.LangPath())
	container.Instance("path.config", path.ConfigPath())
	container.Instance("path.public", path.PublicPath())
	container.Instance("path.storage", path.StoragePath())
	container.Instance("path.database", path.DatabasePath())
	container.Instance("path.resources", path.ResourcePath())
	container.Instance("path.bootstrap", path.BootstrapPath())
}

func (a *Application) Environment() (string, error) {
	if a.Make("env") == "" {
		return "", fmt.Errorf("environment not found")
	}

	return a.Make("env").(string), nil
}

func (a *Application) IsEnvironment(environments ...string) bool {
	for _, environment := range environments {
		if environment == a.Make("env") {
			return true
		}
	}

	return false
}
