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
	logger := getStackWithSingleLogger(t)

	logger.Log(syslog.INFO, "the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - info: the message - `)
}

func TestStackWithMultipleLoggersMustWriteMultipleLogs(t *testing.T) {
	first := loggers.Syslog{Testing: t, Path: testFile, MinLevel: syslog.INFO}
	second := loggers.Syslog{Testing: t, Path: testFile, MinLevel: syslog.INFO}
	allLoggers := map[string]inter.Logger{"first": first, "second": second}
	logger := setUpStack(allLoggers, "first", "second")

	logger.Log(syslog.INFO, "the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 2)
	assert.Contains(t, lines[0][0], ` - - info: the message - `)
	assert.Contains(t, lines[1][0], ` - - info: the message - `)
}

func TestStackWithData(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.LogWith(syslog.INFO, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - info: the message - {"key":"value"}`)
}

func TestStackEmergency(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.Emergency("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - emerg: the message - `)
}

func TestStackEmergencyWithData(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.EmergencyWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - emerg: the message - {"key":12}`)
}

func TestStackAlert(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.Alert("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - alert: the message - `)
}

func TestStackAlertWithData(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.AlertWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - alert: the message - {"key":12}`)
}

func TestStackCritical(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.Critical("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - crit: the message - `)
}

func TestStackCriticalWithData(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.CriticalWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - crit: the message - {"key":12}`)
}

func TestStackError(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.Error("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - err: the message - `)
}

func TestStackErrorWithData(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.ErrorWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - err: the message - {"key":12}`)
}

func TestStackWarning(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.Warning("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - warning: the message - `)
}

func TestStackWarningWithData(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.WarningWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - warning: the message - {"key":12}`)
}

func TestStackNotice(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.Notice("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - notice: the message - `)
}

func TestStackNoticeWithData(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.NoticeWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - notice: the message - {"key":12}`)
}

func TestStackInfo(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.Info("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - info: the message - `)
}

func TestStackInfoWithData(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.InfoWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - info: the message - {"key":12}`)
}

func TestStackDebug(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.Debug("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - debug: the message - `)
}

func TestStackDebugWithData(t *testing.T) {
	logger := getStackWithSingleLogger(t)

	logger.DebugWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` - - debug: the message - {"key":12}`)
}

func getStackWithSingleLogger(t *testing.T) inter.Logger {
	single := loggers.Syslog{Testing: t, Path: testFile, MinLevel: syslog.DEBUG}
	allLoggers := map[string]inter.Logger{"single": single}
	logger := setUpStack(allLoggers, "single")
	return logger
}

func setUpStack(allLoggers map[string]inter.Logger, loggersInStack ...string) inter.Logger {
	setUp()
	app := foundation.NewApp()
	app.Bind("config.Logging.Channels", allLoggers)
	logger := loggers.Stack{Loggers: loggersInStack}
	logger.SetApp(app)
	return logger
}
