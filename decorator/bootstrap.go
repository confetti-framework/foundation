package decorator

import (
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/contract"
)

type BootstrapDecorator struct {
	Bootstraps []contract.Bootstrap
}

func (d BootstrapDecorator) BootstrapWith(app *foundation.Application) *foundation.Application {
	for _, bootstrapper := range d.Bootstraps {
		app = bootstrapper.Bootstrap(app)
	}

	return app
}
