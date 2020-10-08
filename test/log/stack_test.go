package log

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/loggers"
	"github.com/lanvard/syslog"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackWithEmptySlice(t *testing.T) {
	var allLoggers map[string]inter.Logger
	logger := setUpStack(allLoggers)
	logger.Log(syslog.INFO, "the message")

	assert.NoFileExists(t, testFile)
}

func setUpStack(allLoggers map[string]inter.Logger) inter.Logger {
	setUp()
	app := foundation.NewApp()
	app.Instance("config.Logging.Loggers", allLoggers)
	logger := loggers.Stack{}
	logger.SetApp(app)
	return logger
}
