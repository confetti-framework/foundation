package log

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/loggers"
	"github.com/confetti-framework/syslog/log_level"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_log_to_default_channel(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	app.Log().Info("the message")

	require.Len(t, openAndReadFile(testFile), 0)
	require.Len(t, openAndReadFile(testFileSecond), 1)
}

func Test_log_to_non_existing_channel(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	require.PanicsWithError(t, "can not log to channel: channel 'fake' does not exist", func() {
		app.Log("fake").Info("the message")
	})
}

func Test_log_to_multiple_non_existing_channels(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	require.PanicsWithError(t, "can not log to one of the channels: channel 'fake1' does not exist", func() {
		app.Log("fake1", "fake2").Info("the message")
	})
}

func Test_log_to_specific_channel(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	app.Log("first").Info("the message")

	require.Len(t, openAndReadFile(testFile), 1)
	require.Len(t, openAndReadFile(testFileSecond), 0)
}

func Test_log_to_multiple_channels(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	app.Log("first", "second").Info("the message")

	require.Len(t, openAndReadFile(testFile), 1)
	require.Len(t, openAndReadFile(testFileSecond), 1)
}

func getAppWithChannels() inter.App {
	allLoggers := map[string]interface{}{
		"stack":  loggers.Stack{Channels: []string{"first"}},
		"first":  loggers.Syslog{Path: testFile, MinLevel: log_level.INFO},
		"second": loggers.Syslog{Path: testFileSecond, MinLevel: log_level.INFO},
	}

	app := newTestApp()
	app.Bind("config.Logging.Default", "second")
	app.Bind("config.Logging.Channels", allLoggers)

	return app
}
