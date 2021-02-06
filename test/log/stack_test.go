package log

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/loggers"
	"github.com/confetti-framework/syslog/log_level"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_stack_with_empty_slice(t *testing.T) {
	var allLoggers map[string]interface{}
	logger := setUpStack(allLoggers)
	logger.Log(log_level.INFO, "the message")

	require.NoFileExists(t, testFile)
}

func Test_stack_with_one_logger_must_write_one_line(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Log(log_level.INFO, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="info"] the message`)
}

func Test_stack_with_multiple_loggers_must_write_multiple_logs(t *testing.T) {
	first := loggers.Syslog{Path: testFile, MinLevel: log_level.INFO}
	second := loggers.Syslog{Path: testFile, MinLevel: log_level.INFO}
	allLoggers := map[string]interface{}{"first": first, "second": second}
	logger := setUpStack(allLoggers, "first", "second")

	logger.Log(log_level.INFO, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 2)
	require.Contains(t, lines[0][0], ` [level severity="info"] the message`)
	require.Contains(t, lines[1][0], ` [level severity="info"] the message`)
}

func Test_stack_with_data(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(log_level.INFO, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="info"] the message {"key":"value"}`)
}

func Test_stack_emergency(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Log(log_level.EMERGENCY, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="emerg"] the message`)
}

func Test_stack_emergency_with_data(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(log_level.EMERGENCY, "the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="emerg"] the message {"key":12}`)
}

func Test_stack_alert(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Log(log_level.ALERT, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="alert"] the message`)
}

func Test_stack_alert_with_data(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(log_level.ALERT, "the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="alert"] the message {"key":12}`)
}

func Test_stack_critical(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Log(log_level.CRITICAL, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="crit"] the message`)
}

func Test_stack_critical_with_data(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(log_level.CRITICAL, "the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="crit"] the message {"key":12}`)
}

func Test_stack_error(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Log(log_level.ERROR, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="err"] the message`)
}

func Test_stack_error_with_data(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(log_level.ERROR, "the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="err"] the message {"key":12}`)
}

func Test_stack_warning(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Log(log_level.WARNING, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="warning"] the message`)
}

func Test_stack_warning_with_data(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(log_level.WARNING, "the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="warning"] the message {"key":12}`)
}

func Test_stack_notice(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Log(log_level.NOTICE, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="notice"] the message`)
}

func Test_stack_notice_with_data(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(log_level.NOTICE, "the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="notice"] the message {"key":12}`)
}

func Test_stack_info(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Log(log_level.INFO, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="info"] the message`)
}

func Test_stack_info_with_data(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(log_level.INFO, "the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="info"] the message {"key":12}`)
}

func Test_stack_debug(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.Log(log_level.DEBUG, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="debug"] the message`)
}

func Test_stack_debug_with_data(t *testing.T) {
	logger := getStackWithSingleLogger()

	logger.LogWith(log_level.DEBUG, "the message", map[string]int{"key": 12})

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Contains(t, lines[0][0], ` [level severity="debug"] the message {"key":12}`)
}

func getStackWithSingleLogger() inter.Logger {
	single := loggers.Syslog{Path: testFile, MinLevel: log_level.DEBUG}
	allLoggers := map[string]interface{}{"single": single}
	logger := setUpStack(allLoggers, "single")
	return logger
}

func setUpStack(allLoggers map[string]interface{}, loggersInStack ...string) inter.Logger {
	setUp()
	app := newTestApp()
	app.Bind("config.Logging.Channels", allLoggers)
	logger := loggers.Stack{Channels: loggersInStack}
	return logger.SetApp(app)
}
