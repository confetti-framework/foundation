package console

import "github.com/confetti-framework/contract/inter"

var FoundationCommands = []inter.Command{
	AppInfo{},
	AppServe{},
	LogClear{},
	RouteList{},
}
