package log

import (
	"github.com/lanvard/contract/inter"
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
	logger := getStackWithSingleLogger()

	logger.Log(syslog.INFO, "the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="info"] `)
}

func TestStackWithMultipleLoggersMustWriteMultipleLogs(t *testing.T) {
	first := loggers.Syslog{Path: testFile, MinLevel: syslog.INFO}
	second := loggers.Syslog{Path: testFile, MinLevel: syslog.INFO}
	allLoggers := map[string]inter.Logger{"first": first, "second": second}
	logger := setUpStack(allLoggers, "first", "second")

	logger.Log(syslog.INFO, "the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 2)
	assert.Contains(t, lines[0][0], ` the message [level severity="info"] `)
	assert.Contains(t, lines[1][0], ` the message [level severity="info"] `)
}

func TestStackWithData(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(syslog.INFO, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="info"] {"key":"value"}`)
}

func TestStackEmergency(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Emergency("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="emerg"] `)
}

func TestStackEmergencyWithData(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.EmergencyWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="emerg"] {"key":12}`)
}

func TestStackAlert(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Alert("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="alert"] `)
}

func TestStackAlertWithData(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.AlertWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="alert"] {"key":12}`)
}

func TestStackCritical(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Critical("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="crit"] `)
}

func TestStackCriticalWithData(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.CriticalWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="crit"] {"key":12}`)
}

func TestStackError(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Error("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="err"] `)
}

func TestStackErrorWithData(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.ErrorWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="err"] {"key":12}`)
}

func TestStackWarning(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Warning("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="warning"] `)
}

func TestStackWarningWithData(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.WarningWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="warning"] {"key":12}`)
}

func TestStackNotice(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Notice("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="notice"] `)
}

func TestStackNoticeWithData(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.NoticeWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="notice"] {"key":12}`)
}

func TestStackInfo(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Info("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="info"] `)
}

func TestStackInfoWithData(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.InfoWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="info"] {"key":12}`)
}

func TestStackDebug(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Debug("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="debug"] `)
}

func TestStackDebugWithData(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.DebugWith("the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
	assert.Contains(t, lines[0][0], ` the message [level severity="debug"] {"key":12}`)
}

func getStackWithSingleLogger() inter.Logger {
	single := loggers.Syslog{Path: testFile, MinLevel: syslog.DEBUG}
	allLoggers := map[string]inter.Logger{"single": single}
	logger := setUpStack(allLoggers, "single")
	return logger
}

func setUpStack(allLoggers map[string]inter.Logger, loggersInStack ...string) inter.Logger {
	setUp()
	app := newTestApp()
	app.Bind("config.Logging.Channels", allLoggers)
	logger := loggers.Stack{Channels: loggersInStack}
	return logger.SetApp(app)
}
