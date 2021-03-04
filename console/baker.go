package console

import (
	"github.com/confetti-framework/baker"
	"github.com/confetti-framework/contract/inter"
	"os"
)

type Baker struct{}

func (b Baker) Name() string {
	return "baker"
}

func (b Baker) Description() string {
	return "Interact with your application."
}

func (b Baker) Handle(c inter.Cli) inter.ExitCode {
	err := gore.New(
		gore.AutoImport(true),
		gore.OutWriter(c.Writer()),
		gore.ErrWriter(os.Stderr),
	).Run()

	if err != nil {
		c.Error("Baker: %s\n", err)
		return inter.Failure
	}
	return inter.Success
}
