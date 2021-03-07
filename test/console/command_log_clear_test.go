package console

import (
	"bytes"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation/console"
	"github.com/confetti-framework/foundation/loggers"
	"github.com/confetti-framework/syslog/log_level"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"testing"
)

func Test_log_clear_get_name(t *testing.T) {
	assert.Equal(t, "log:clear", console.LogClear{}.Name())
}

func Test_log_clear_get_description(t *testing.T) {
	assert.Equal(t, "Clear the log files as indicated in the configuration.", console.LogClear{}.Description())
}

func Test_log_clear_without_loggers(t *testing.T) {
	writer, app := setUp()
	var writerErr bytes.Buffer

	app.Bind("config.App.OsArgs", []interface{}{"/main", "log:clear"})
	app.Bind("config.Logging.Channels", map[string]interface{}{})

	code := console.Kernel{
		App:       app,
		Writer:    &writer,
		WriterErr: &writerErr,
		Commands:  []inter.Command{console.LogClear{}},
	}.Handle()

	require.Equal(t, inter.Success, code)
	require.Contains(t, writerErr.String(), `No files to clear. No loggers found`)
	require.Contains(t, writer.String(), `Done`)
}

func Test_log_clear_with_recent_log_file(t *testing.T) {
	// Given
	writer, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/main", "log:clear"})

	dir, err := ioutil.TempDir("", "log_clear_without_loggers_")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	allLoggers := map[string]interface{}{
		"test_logger": loggers.Syslog{Path: dir + "/{yyyy-mm-dd}_test.log", MinLevel: log_level.INFO, MaxFiles: 1},
	}
	app.Bind("config.Logging.Default", "test_logger")
	app.Bind("config.Logging.Channels", allLoggers)
	app.Log().Info("test message")

	// When
	code := console.Kernel{
		App:      app,
		Writer:   &writer,
		Commands: []inter.Command{console.LogClear{}},
	}.Handle()

	// Then
	require.Equal(t, inter.Success, code)
	require.Contains(t, TrimDoubleSpaces(writer.String()), `Done`)

	readDir, err := ioutil.ReadDir(dir)
	require.Nil(t, err)
	require.Len(t, readDir, 1)
}

//goland:noinspection GoNilness
func Test_log_clear_with_old_logger_file(t *testing.T) {
	// Given
	writer, app := setUp()
	app.Bind("config.App.OsArgs", []interface{}{"/main", "log:clear"})

	dir, err := ioutil.TempDir("", "log_clear_without_loggers_")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	allLoggers := map[string]interface{}{
		"test_logger": loggers.Syslog{Path: dir + "/{yyyy-mm-dd}_test.log", MinLevel: log_level.INFO, MaxFiles: 1},
	}
	app.Bind("config.Logging.Default", "test_logger")
	app.Bind("config.Logging.Channels", allLoggers)

	file, err := os.Create(dir + "/1991-01-02_test.log")
	defer file.Close()

	file, err = os.Create(dir + "/1991-01-01_test.log")
	defer file.Close()

	// When
	code := console.Kernel{
		App:      app,
		Writer:   &writer,
		Commands: []inter.Command{console.LogClear{}},
	}.Handle()

	// Then
	require.Equal(t, inter.Success, code)
	require.Contains(t, TrimDoubleSpaces(writer.String()), `Done`)

	readDir, err := ioutil.ReadDir(dir)
	require.Nil(t, err)
	require.Len(t, readDir, 1)
	require.Regexp(
		t,
		`Files cleaned for channel: test_logger.*\n`,
		TrimDoubleSpaces(writer.String()),
	)
}

func TrimDoubleSpaces(value string) string {
	// Replace double spaces
	regex := regexp.MustCompile(` {2,}`)
	value = regex.ReplaceAllString(strings.Trim(value, " "), " ")

	// replace newline with only one space
	regex = regexp.MustCompile(` \n`)
	value = regex.ReplaceAllString(value, "\n")

	return value
}
