package foundation

import (
	"fmt"
	"github.com/lanvard/contract/inter"
	"lanvard/config"
)

type Application struct {
	// The service container
	container *Container

	// Indicates if the application has been bootstrapped before.
	HasBeenBootstrapped bool
}

// Get the service container
func (a *Application) Container() *Container {
	return a.container
}

// Set the service container
func (a *Application) SetContainer(container *Container) {
	a.container = container
}

// Resolve the given type from the container.
func (a Application) Make(abstract interface{}) interface{} {
	return a.container.Make(abstract)
}

// Bind all of the application paths in the container.
func (a *Application) BindPathsInContainer(path inter.BasePath) {
	a.container.Instance("path.app", path.AppPath())
	a.container.Instance("path.base", path.BasePath())
	a.container.Instance("path.lang", path.LangPath())
	a.container.Instance("path.config", path.ConfigPath())
	a.container.Instance("path.public", path.PublicPath())
	a.container.Instance("path.storage", path.StoragePath())
	a.container.Instance("path.database", path.DatabasePath())
	a.container.Instance("path.resources", path.ResourcePath())
	a.container.Instance("path.bootstrap", path.BootstrapPath())
}

func (a *Application) Environment() (string, error) {
	if config.App.Env == "" {
		return "", fmt.Errorf("environment not found")
	}

	return config.App.Env, nil
}

func (a *Application) IsEnvironment(environments ...string) bool {
	for _, environment := range environments {
		if environment == config.App.Env {
			return true
		}
	}

	return false
}
