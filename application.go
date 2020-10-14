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
	(*a.container).Singleton(abstract, concrete)
}

// Resolve the given type from the container.
func (a *Application) Make(abstract interface{}) interface{} {
	return (*a.container).Make(abstract)
}

func (a *Application) Bind(abstract interface{}, concrete interface{}) {
	(*a.container).Bind(abstract, concrete)
}

func (a *Application) Instance(abstract interface{}) interface{} {
	return (*a.container).Instance(abstract)
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

func (a *Application) Log() inter.Logger {
	key := a.Make("config.Logging.Default").(string)
	logger := a.Make("config.Logging.Channels").(map[string]interface{})[key].(inter.Logger)

	logger = logger.SetApp(a)

	return logger
}
