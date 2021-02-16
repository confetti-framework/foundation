package console

import (
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console/service"
	"io"
	"os"
)

type Kernel struct {
	App      inter.App
	Commands []inter.Command
	Output   io.Writer
}

func (k Kernel) Handle() inter.ExitCode {
	if k.Output == nil {
		k.Output = os.Stdout
	}

	args := service.GetArgs(k.App)
	if len(args) > 1 {
		code := service.DispatchCommands(k.App, k.Output, k.Commands)
		if code != inter.Help {
			// todo test help page
			_, _ = fmt.Fprintln(k.Output, "Done")
			return code
		}
	}

	return service.RenderIndex(k.Output, k.Commands)
}
