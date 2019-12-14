package decorator

import (
	"github.com/lanvard/contract/decorator"
	"github.com/lanvard/foundation"
)

type BootstrapDecorator struct {
	Bootstraps []decorator.Bootstrap
}

func (d BootstrapDecorator) BootstrapWith(app foundation.Application) foundation.Application {
	for _, bootstrapper := range d.Bootstraps {
		app = bootstrapper.Bootstrap(app)
	}

	return app
}
