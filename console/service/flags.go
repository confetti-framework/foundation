package service

import (
	"flag"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console/flag_type"
	"github.com/confetti-framework/support"
	"reflect"
)

func ActualCommandName(appArgs []string) string {
	if len(appArgs) > 1 {
		return appArgs[1]
	}
	return ""
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

func registerFlags(set *flag.FlagSet, options []Field, getters []flag.Getter) {
	for _, option := range options {
		// Clone the getters so that the value of one flag cannot be set to another.
		result := cloneGetters(getters)

		registerFlag(set, option, option.Tag.Get("short"), "", result)
		registerFlag(set, option, option.Tag.Get("flag"), option.Tag.Get("description"), result)
	}
	set.Var(new(flag_type.String), "env-file", "Can be used with any command to run the command with a defined environment file.")
}

func registerFlag(set *flag.FlagSet, option Field, flag string, description string, getters []flag.Getter) {
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
		if getterName == support.Name(option.Value) {
			set.Var(getter, flag, description)
		}
	}
}

//goland:noinspection GoPreferNilSlice
func cloneGetters(getters []flag.Getter) []flag.Getter {
	result := []flag.Getter{}
	for _, oldObj := range getters {
		indirect := reflect.Indirect(reflect.ValueOf(oldObj))
		newIndirect := reflect.New(indirect.Type())
		newIndirect.Elem().Set(reflect.ValueOf(indirect.Interface()))
		newNamed := newIndirect.Interface()
		result = append(result, newNamed.(flag.Getter))
	}

	return result
}
