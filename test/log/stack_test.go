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

func TestStackWithOneLoggerMustWriteOneLine(t *testing.T) {
	single := loggers.Syslog{Testing: t, Path: testFile}
	allLoggers := map[string]inter.Logger{"single": single}
	logger := setUpStack(allLoggers, "single")

	logger.Log(syslog.INFO, "the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - info: the message - `)
}

func setUpStack(allLoggers map[string]inter.Logger, loggersInStack ...string) inter.Logger {
	setUp()
	app := foundation.NewApp()
	app.Instance("config.Logging.Loggers", allLoggers)
	logger := loggers.Stack{Loggers: loggersInStack}
	logger.SetApp(app)
	return logger
}
