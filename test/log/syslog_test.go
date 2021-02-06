package log

import (
	"encoding/csv"
	"fmt"
	"github.com/confetti-framework/foundation/loggers"
	"github.com/confetti-framework/syslog"
	"github.com/confetti-framework/syslog/log_level"
	"github.com/stretchr/testify/require"
	"io"
	"log"
	"os"
	"testing"
)

const testDir = "/tmp/error_tests/"
const testFile = testDir + "log_first.log"
const testFileSecond = testDir + "log_second.log"

var structMock = struct{ FirstLevel interface{} }{
	struct{ SecondLevel string }{
		"ceiling",
	},
}

func Test_create_log_file_if_not_exists(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.INFO, "some content")

	require.FileExists(t, testFile)
}

func Test_get_add_one_error_line(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.INFO, "error line 1")

	require.Len(t, openAndReadFile(testFile), 1)
}

func Test_get_add_two_error_lines(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.INFO, "error line 1")
	logger.Log(log_level.INFO, "error line 2")

	require.Len(t, openAndReadFile(testFile), 2)
}

func Test_log_with_level(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.INFO, "message")

	lines := openAndReadFile(testFile)
	require.Regexp(t, "^<6.*\"info\"", lines[0][0])
}

func Test_content_of_error(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.INFO, "error line 1")
	logger.Log(log_level.INFO, "error line 2")

	lines := openAndReadFile(testFile)
	require.Regexp(t, `^<6>1 \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.* \[level severity="info"\] error line 1`, lines[0][0])
	require.Regexp(t, `^<6>1 \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.* \[level severity="info"\] error line 2`, lines[1][0])
}

func Test_log_arguments(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.INFO, "name=%s", "Ron")

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], "name=Ron")
}

func Test_log_with_string(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.INFO, "the message", "string data")

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="info"\] the message string data$`, lines[0][0])
}

func Test_log_with_map(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.INFO, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], ` [level severity="info"] the message {"key":"value"}`)
}

func Test_log_with_struct(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.INFO, "the message", structMock)

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], `- [level severity="info"] the message {"FirstLevel":{"SecondLevel":"ceiling"}}`)
}

func Test_log_levels(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.INFO, "the message", structMock)

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], `- [level severity="info"] the message {"FirstLevel":{"SecondLevel":"ceiling"}}`)
}

func Test_log_type(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1).Group("external")

	logger.LogWith(log_level.INFO, "the message", structMock)

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], ` external [level severity="info"] the message {"FirstLevel":{"SecondLevel":"ceiling"}}`)
}

func Test_log_with_structured_data(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)
	data := syslog.StructuredData{
		"firstLevel": {"secondLevel": "the value"},
	}

	logger.LogWith(log_level.INFO, "the message", data)

	lines := openAndReadFile(testFile)

	require.Contains(t, lines[0][0], ` [firstLevel secondLevel="the value"][level severity="info"] the message`)
}

func Test_log_emergency(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.EMERGENCY, "the message")

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="emerg"\] the message $`, lines[0][0])
}

func Test_log_emergency_arguments(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.EMERGENCY, "name=%s", "Ron")

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], "name=Ron")
}

func Test_log_emergency_with_data(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.EMERGENCY, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="emerg"\] the message {"key":"value"}$`, lines[0][0])
}

func Test_log_alert(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.ALERT, "the message")

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="alert"\] the message $`, lines[0][0])
}

func Test_log_alert_arguments(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.ALERT, "name=%s", "Ron")

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], "name=Ron")
}

func Test_log_alert_with_data(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.ALERT, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="alert"\] the message {"key":"value"}$`, lines[0][0])
}

func Test_log_critical(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.CRITICAL, "the message")

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="crit"\] the message $`, lines[0][0])
}

func Test_log_critical_arguments(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.CRITICAL, "name=%s", "Ron")

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], "name=Ron")
}

func Test_log_critical_with_data(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.CRITICAL, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="crit"\] the message {"key":"value"}$`, lines[0][0])
}

func Test_log_error(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.ERROR, "the message")

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="err"\] the message $`, lines[0][0])
}

func Test_log_error_arguments(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.ERROR, "name=%s", "Ron")

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], "name=Ron")
}

func Test_log_error_with_data(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.ERROR, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="err"\] the message {"key":"value"}$`, lines[0][0])
}

func Test_log_warning(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.WARNING, "the message")

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="warning"\] the message $`, lines[0][0])
}

func Test_log_warning_arguments(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.WARNING, "name=%s", "Ron")

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], "name=Ron")
}

func Test_log_warning_with_data(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.WARNING, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="warning"\] the message {"key":"value"}$`, lines[0][0])
}

func Test_log_notice(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.NOTICE, "the message")

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="notice"\] the message $`, lines[0][0])
}

func Test_log_notice_arguments(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.NOTICE, "name=%s", "Ron")

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], "name=Ron")
}

func Test_log_notice_with_data(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.NOTICE, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="notice"\] the message {"key":"value"}$`, lines[0][0])
}

func Test_log_info(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.INFO, "the message")

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="info"\] the message $`, lines[0][0])
}

func Test_log_info_arguments(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.INFO, "name=%s", "Ron")

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], "name=Ron")
}

func Test_log_info_with_data(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.INFO, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="info"\] the message {"key":"value"}$`, lines[0][0])
}

func Test_log_debug(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.DEBUG, "the message")

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="debug"\] the message $`, lines[0][0])
}

func Test_log_debug_arguments(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.DEBUG, "name=%s", "Ron")

	lines := openAndReadFile(testFile)
	require.Contains(t, lines[0][0], "name=Ron")
}

func Test_log_debug_with_data(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(log_level.DEBUG, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	require.Regexp(t, ` \[level severity="debug"\] the message {"key":"value"}$`, lines[0][0])
}

func Test_log_with_min_level(t *testing.T) {
	setUp()
	logger := loggers.Syslog{Path: testFile, MinLevel: log_level.INFO}

	logger.Log(log_level.DEBUG, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 0)
}

func Test_log_same_level_as_min_level(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(log_level.INFO, "the message")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
}

func Test_log_with_empty_group(t *testing.T) {
	setUp()
	app := getAppWithChannels()

	app.Log().Group("").Info("response: ok")

	lines := openAndReadFile(testFileSecond)
	require.Len(t, lines, 1)
	require.Regexp(t, ` - \[level severity="info"\] response: ok $`, lines[0][0])
}

func Test_log_with_group_on_stack_logger(t *testing.T) {
	setUp()
	app := getAppWithChannels()
	app.Bind("config.Logging.Default", "stack")

	app.Log().Group("external").Info("response: ok")

	lines := openAndReadFile(testFile)
	require.Len(t, lines, 1)
	require.Regexp(t, ` external \[level severity="info"\] response: ok $`, lines[0][0])
}

func setUp() {
	// Remove old test files
	_ = os.RemoveAll(testDir)

	// create empty test dir
	err := os.MkdirAll(testDir, 0755)
	if err != nil {
		panic(err)
	}
}

func openAndReadFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	defer func() { _ = file.Close() }()
	if err != nil {
		return [][]string{}
	}
	lines, err := readFile(file)
	if err != nil {
		fmt.Printf("Failed to read file: %s\n", fileName)
	}
	return lines
}

func readFile(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	r.LazyQuotes = true
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return lines, err
}
