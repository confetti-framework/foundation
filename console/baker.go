package console

import (
	"fmt"
	gore "github.com/confetti-framework/baker"
	"github.com/confetti-framework/contract/inter"
	"io"
	"os"
)

type Baker struct{}

func (b Baker) Name() string {
	return "baker"
}

func (b Baker) Description() string {
	return "Interact with your application."
}

func (b Baker) Handle(_ inter.App, writer io.Writer) inter.ExitCode {
	err := gore.New(
		gore.AutoImport(true),
		gore.OutWriter(writer),
		gore.ErrWriter(os.Stderr),
	).Run()

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Baker: %s\n", err)
		return inter.Failure
	}
	return inter.Success
}
