package service

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"io"
	"reflect"
	"strings"
)

func DispatchCommands(app inter.App, output io.Writer, commands []inter.Command) inter.ExitCode {
	for _, command := range commands {
		code := HandleCommand(app, output, command)
		if code != inter.Continue {
			return code
		}
	}

	return inter.Help
}

func HandleCommand(app inter.App, output io.Writer, command inter.Command) inter.ExitCode {

	// gc := StructWithOptionBool{DryRun: false}

	flagSet := flag.NewFlagSet(command.Name(), flag.ExitOnError)
	flagSet.SetOutput(output)
	appArgs := GetArgs(app)

	if appArgs[1] == command.Name() {
		options := GetOptions(command)
		option := options[0]
		// todo set all flags

		_ = flagSet.Parse(appArgs[1:])



		appArg := strings.TrimLeft(appArgs[2], "-")
		if option.Tag.Get("flag") == appArg {

			// v is the interface{}
			v := reflect.ValueOf(&command).Elem()

			// Allocate a temporary variable with type of the struct.
			//    v.Elem() is the value contained in the interface.
			tmp := reflect.New(v.Elem().Type()).Elem()

			// Copy the struct value contained in interface to
			// the temporary variable.
			tmp.Set(v.Elem())

			// Set the field.
			tmp.Field(option.Number).SetBool(true)

			// Set the interface to the modified struct value.
			v.Set(tmp)
		}

		return command.Handle(app, output)
	}

	return inter.Continue
}

// GetArgs converts arguments from config []interface{} to []string
func GetArgs(app inter.App) []string {
	var result []string
	argsRaw := app.Make("config.App.OsArgs")
	for _, argRaw := range argsRaw.([]interface{}) {
		result = append(result, argRaw.(string))
	}

	return result
}
