package console

import (
	"github.com/confetti-framework/contract/inter"
)

type LogClear struct {}

func (l LogClear) Name() string {
	return "log:clear"
}

func (l LogClear) Description() string {
	return "Clear the log files as indicated in the configuration."
}

func (l LogClear) Handle(c inter.Cli) inter.ExitCode {
	channelsRaw, err := c.App().MakeE("config.Logging.Channels")
	loggers := channelsRaw.(map[string]interface{})
	if err != nil || len(loggers) == 0 {
		c.Error("No files to clear. No loggers found")
	}

	for channel, logger := range loggers {
		if logger.(inter.Logger).Clear() {
			c.Info("Files cleaned for channel: %s", channel)
		}
	}

	c.Info("Done")

	return inter.Success
}
