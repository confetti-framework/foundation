package foundation

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/foundation/loggers"
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

// GetE the service container
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

// Make the given type from the container.
func (a *Application) Make(abstract interface{}) interface{} {
	return (*a.container).Make(abstract)
}

// Make or give an error by the given type from the container.
func (a *Application) MakeE(abstract interface{}) (interface{}, error) {
	return (*a.container).MakeE(abstract)
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

func (a *Application) Log(channels ...string) inter.Logger {
	var logger inter.Logger

	// If no channels are specified, take the default
	if len(channels) == 0 {
		defaultLogger, err := a.MakeE("config.Logging.Default")
		if err != nil {
			panic("no default logger found in config.Logging.Default")
		}
		channels = []string{defaultLogger.(string)}
	}

	//  If multiple loggers are present, wrap them in a stack
	if len(channels) > 1 {
		logger = loggers.Stack{Channels: channels}
	} else {
		logger = a.getLoggerByChannel(channels[0])
	}

	return logger.SetApp(a)
}

func (a *Application) getLoggerByChannel(channel string) inter.Logger {
	all := a.Make("config.Logging.Channels").(map[string]interface{})
	rawLogger, ok := all[channel]
	if !ok {
		panic(errors.New("can not log to channel. Channel '" + channel + "' does not exist"))
	}

	return rawLogger.(inter.Logger)
}
