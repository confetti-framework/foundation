package providers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
)

type DatabaseServiceProvider struct {
	Connections map[string]inter.Connection
}

func (c DatabaseServiceProvider) Boot(container inter.Container) inter.Container {
	for name, connection := range c.Connections {
		err := connection.Open()
		if err != nil {
			panic(err)
		}

		if err = connection.Pool().Ping(); err != nil {
			panic(errors.Wrap(err, "can't connect with connection " + name))
		}
	}

	container.Bind("open_connections", c.Connections)

	return container
}
