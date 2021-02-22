package console

import (
	"flag"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console/flag_type"
	"github.com/confetti-framework/foundation/console/service"
	"io"
	"os"
)

// These values must be in a callback so that they are new
// for each field. If there are multiple fields with the same
// type, they would use the same pointer.
var flagGetters = func() []flag.Getter {
	return []flag.Getter{
		new(flag_type.Bool),
		new(flag_type.String),
		new(flag_type.Int),
		new(flag_type.Int64),
		new(flag_type.Uint),
		new(flag_type.Uint64),
		new(flag_type.Duration),
	}
}

type Kernel struct {
	App           inter.App
	Commands      []inter.Command
	Output        io.Writer
	FlagProviders []func() []flag.Getter
}

func (k Kernel) Handle() inter.ExitCode {
	if k.Output == nil {
		k.Output = os.Stdout
	}

	k.FlagProviders = append(k.FlagProviders, flagGetters)

	code := service.DispatchCommands(k.App, k.Output, k.Commands, k.FlagProviders)
	if code != inter.Help {
		_, _ = fmt.Fprintln(k.Output, "Done")
		return code
	}

	return service.RenderIndex(k.App, k.Output, k.Commands)
}
