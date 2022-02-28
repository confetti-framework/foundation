package console

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console/facade"
	"github.com/confetti-framework/foundation/console/flag_type"
	"github.com/confetti-framework/foundation/console/service"
	"io"
)

// These values must be in a callback so that they are new
// for each field. If there are multiple fields with the same
// type, they would use the same pointer.
var flagGetters = []flag.Getter{
	new(flag_type.Bool),
	new(flag_type.String),
	new(flag_type.Int),
	new(flag_type.Int64),
	new(flag_type.Uint),
	new(flag_type.Uint64),
	new(flag_type.Duration),
}

type Kernel struct {
	App           inter.App
	Writer        io.Writer
	WriterErr     io.Writer
	Commands      []inter.Command
	FlagProviders []func() []flag.Getter // Deprecated: FlagProviders is deprecated and will be removed in version 1.0. Use Getters instead.
	Getters       []flag.Getter          // This list include custom flag.Getters, you can create custom types to cast flags from the command to a value.
}

func (k Kernel) Handle() inter.ExitCode {
	getters := append(k.Getters, flagGetters...)

	// FlagProviders is deprecated and will be removed in version 1.0. Use Getters instead.
	for _, provider := range k.FlagProviders {
		getters = append(getters, provider()...)
	}

	cli := facade.NewCli(k.App, k.Writer, k.WriterErr)
	code := service.DispatchCommands(cli, k.Commands, getters)
	if code != inter.Index {
		return code
	}

	return service.RenderIndex(cli, k.Commands)
}

func (k Kernel) GetCommands() []inter.Command {
	return k.Commands
}
