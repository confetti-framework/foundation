package decorator

import (
	"github.com/lanvard/contract/inter"
)

type BootstrapDecorator struct {
	Bootstraps []inter.Bootstrap
}

func (d BootstrapDecorator) BootstrapWith(app inter.App) inter.App {
	for _, bootstrapper := range d.Bootstraps {
		app = bootstrapper.Bootstrap(app)
	}

	return app
}
