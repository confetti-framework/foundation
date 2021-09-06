package console

import (
	"github.com/confetti-framework/contract/inter"
	"runtime"
)

// AppInfo shows the info of the application.
type AppInfo struct {
}

// Name of the command
func (s AppInfo) Name() string {
	return "app:info"
}

// Description of the command
func (s AppInfo) Description() string {
	return "Show the info of the application."
}

// Handle contains the logic of the command
func (s AppInfo) Handle(c inter.Cli) inter.ExitCode {

	t := c.Table()
	t.AppendRow([]interface{}{"Golang Version", runtime.Version()})
	t.AppendRow([]interface{}{"Number of CPUs", runtime.NumCPU()})

	t.Render()

	return inter.Success
}
