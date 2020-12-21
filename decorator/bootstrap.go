package decorator

import (
	"github.com/confetti-framework/contract/inter"
)

type BootstrapDecorator struct {
	Bootstraps []inter.Bootstrap
}

func (d BootstrapDecorator) BootstrapWith(container inter.Container) inter.Container {
	for _, bootstrapper := range d.Bootstraps {
		container = bootstrapper.Bootstrap(container)
	}

	return container
}
