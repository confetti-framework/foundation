package console

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console/flag_type"
	"github.com/confetti-framework/foundation/console/service"
	"io"
	"os"
	"strconv"
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

	args := k.getArgs()

	if len(args) > 1 {

		flagSet := flag.NewFlagSet("serve", flag.ExitOnError)
		flagSet.SetOutput(k.Output)

		// for _, option := range options {
		// flagValues := k.App.Make((*map[string]flag.Value)(nil)).(map[string]flag.Value)
		// flagValue := flagValues["string"]
		var result flag_type.String
		// firstFlag := strings.TrimLeft(option.Tag.Get("flag"), "-")
		// option.Tag.Get("description")
		flagSet.Var(&result, "port", "this is a description")
		if args[1] == "serve" {
			_ = flagSet.Parse(args[2:])
			port, _ := strconv.Atoi(string(result))
			AppServe{Port: port}.Handle(k.App, k.Output)
		}
	}

	return service.RenderIndex(k.Output, k.Commands)
}

// Convert arguments from config []interface{} to []string
func (k Kernel) getArgs() []string {
	var result []string
	argsRaw := k.App.Make("config.App.OsArgs")
	for _, argRaw := range argsRaw.([]interface{}) {
		result = append(result, argRaw.(string))
	}

	return result
}
