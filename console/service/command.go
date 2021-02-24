package service

import (
	"flag"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"io"
	"reflect"
	"strings"
)

func DispatchCommands(
	app inter.App,
	w io.Writer,
	commands []inter.Command,
	flagProviders []func() []flag.Getter,
) inter.ExitCode {
	for _, command := range commands {
		code := handleCommand(app, w, command, flagProviders)
		if code != inter.Continue {
			return code
		}
	}

	return inter.Help
}

func handleCommand(
	app inter.App,
	w io.Writer,
	command inter.Command,
	flagProviders []func() []flag.Getter,
) inter.ExitCode {
	set := flag.NewFlagSet(command.Name(), flag.ContinueOnError)
	set.SetOutput(w)

	actualArgs := actualArgs(app)

	if ActualCommandName(actualArgs) == command.Name() {
		options := GetOptions(command)
		set.Usage = helpFormat(w, command, options)

		registerFlags(set, options, flagProviders)
		code := parse(set, actualArgs)
		if code != inter.Continue {
			return code
		}
		setValuesInCommand(&command, set, options, actualArgs)

		return command.Handle(app, w)
	}

	return inter.Continue
}

func setValuesInCommand(command *inter.Command, set *flag.FlagSet, options []Field, actualArgs []string) {
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

func flagGettersByProviders(providers []func() []flag.Getter) []flag.Getter {
	var result []flag.Getter
	for _, provider := range providers {
		result = append(result, provider()...)
	}
	return result
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
