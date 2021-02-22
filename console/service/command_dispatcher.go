package service

import (
	"flag"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/support"
	"io"
	"reflect"
	"strings"
)

func DispatchCommands(
	app inter.App,
	output io.Writer,
	commands []inter.Command,
	flagProviders []func() []flag.Getter,
) inter.ExitCode {
	for _, command := range commands {
		code := handleCommand(app, output, command, flagProviders)
		if code != inter.Continue {
			return code
		}
	}

	return inter.Help
}

func handleCommand(
	app inter.App,
	output io.Writer,
	command inter.Command,
	flagProviders []func() []flag.Getter,
) inter.ExitCode {
	flagSet := flag.NewFlagSet(command.Name(), flag.ContinueOnError)
	flagSet.SetOutput(output)
	actualArgs := actualArgs(app)

	if actualCommandName(actualArgs) == command.Name() {
		options := GetOptions(command)

		registerOptions(flagSet, options, flagProviders)
		code := parse(flagSet, actualArgs)
		if code != inter.Continue {
			return code
		}
		setValuesInCommand(&command, flagSet, options, actualArgs)

		return command.Handle(app, output)
	}

	return inter.Continue
}

func setValuesInCommand(command *inter.Command, set *flag.FlagSet, options []ParsedOption, actualArgs []string) {
	for _, option := range options {
		for _, actual := range actualArgs[2:] {
			actual := strings.TrimLeft(actual, "-")
			if actual == option.Tag.Get("short") {
				setValueByFlag(command, option.Number, getValue(set, option.Tag.Get("short")))
			} else if actual == option.Tag.Get("flag") {
				setValueByFlag(command, option.Number, getValue(set, option.Tag.Get("flag")))
			}
		}
	}
}

func getValue(set *flag.FlagSet, key string) flag.Getter {
	rawValue, ok := set.Lookup(key).Value.(flag.Getter)
	if ok == false {
		panic(fmt.Sprintf("Can't get value from flag type. Flag for key %s does not implement `flag.Getter`", key))
	}
	return rawValue
}

func actualCommandName(appArgs []string) string {
	if len(appArgs) > 1 {
		return appArgs[1]
	}
	return ""
}

func parse(flagSet *flag.FlagSet, appArgs []string) inter.ExitCode {
	var err error = nil
	if len(appArgs) > 1 {
		err = flagSet.Parse(appArgs[2:])
	}
	if err != nil {
		return inter.Failure
	}
	return inter.Continue
}

func registerOptions(set *flag.FlagSet, options []ParsedOption, flagProviders []func() []flag.Getter) {
	for _, option := range options {
		fg := flagGettersByProviders(flagProviders)
		registerFlag(set, option, option.Tag.Get("short"), "", fg)
		registerFlag(set, option, option.Tag.Get("flag"), option.Tag.Get("description"), fg)
	}
}

func flagGettersByProviders(providers []func() []flag.Getter) []flag.Getter {
	var result []flag.Getter
	for _, provider := range providers {
		result = append(result, provider()...)
	}
	return result
}

func registerFlag(set *flag.FlagSet, option ParsedOption, flag string, description string, getters []flag.Getter) {
	// Check if the flag is already present.
	if set.Lookup(flag) != nil || flag == "" {
		return
	}

	// There is no need to add a - to a tag. This will be added automatically.
	if flag[0] == '-' {
		panic("field with tag `" + flag + "` starts with a -. That's not allowed.")
	}

	for _, getter := range getters {
		getterName := support.Name(getter.Get())
		valueName := support.Name(option.Value)
		if getterName == valueName {
			set.Var(getter, flag, description)
		}
	}
}

func setValueByFlag(command *inter.Command, i int, rawValue flag.Getter) {
	// v is the interface{}
	v := reflect.ValueOf(command).Elem()

	// Allocate a temporary variable with type of the struct.
	// v.Elem() is the value contained in the interface.
	tmp := reflect.New(v.Elem().Type()).Elem()

	// Copy the struct value contained in interface to
	// the temporary variable.
	tmp.Set(v.Elem())

	// Set the field.
	value := reflect.ValueOf(rawValue.Get())
	tmp.Field(i).Set(value)

	// Set the interface to the modified struct value.
	v.Set(tmp)
}

// actualArgs converts arguments from config []interface{} to []string
func actualArgs(app inter.App) []string {
	var result []string
	argsRaw := app.Make("config.App.OsArgs")
	for _, argRaw := range argsRaw.([]interface{}) {
		result = append(result, argRaw.(string))
	}

	return result
}
