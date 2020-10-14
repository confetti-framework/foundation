package log

import (
	"encoding/csv"
	"fmt"
	"github.com/lanvard/foundation/loggers"
	"github.com/lanvard/syslog"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"testing"
)

const testDir = "/tmp/error_tests/"
const testFile = testDir + "log_test.log"

var structMock = struct{ FirstLevel interface{} }{
	struct{ SecondLevel string }{
		"ceiling",
	},
}

func TestCreateLogFileIfNotExists(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(syslog.INFO, "some content")

	assert.FileExists(t, testFile)
}

func TestGetAddOneErrorLine(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(syslog.INFO, "error line 1")

	assert.Len(t, openAndReadFile(testFile), 1)
}

func TestGetAddTwoErrorLines(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(syslog.INFO, "error line 1")
	logger.Log(syslog.INFO, "error line 2")

	assert.Len(t, openAndReadFile(testFile), 2)
}

func TestLogWithLevel(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(syslog.INFO, "message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, "^<6.*\"info\"", lines[0][0])
}

func TestContentOfError(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Log(syslog.INFO, "error line 1")
	logger.Log(syslog.INFO, "error line 2")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `^<6>1 \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.* error line 1 \[level severity="info"\]`, lines[0][0])
	assert.Regexp(t, `^<6>1 \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.* error line 2 \[level severity="info"\]`, lines[1][0])
}

func TestLogWithString(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(syslog.INFO, "the message", "string data")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` the message \[level severity="info"\] string data$`, lines[0][0])
}

func TestLogWithMap(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(syslog.INFO, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Contains(t, lines[0][0], `the message [level severity="info"] {"key":"value"}`)
}

func TestLogWithStruct(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(syslog.INFO, "the message", structMock)

	lines := openAndReadFile(testFile)
	assert.Contains(t, lines[0][0], `the message [level severity="info"] {"FirstLevel":{"SecondLevel":"ceiling"}}`)
}

func TestLogLevels(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.LogWith(syslog.INFO, "the message", structMock)

	lines := openAndReadFile(testFile)
	assert.Contains(t, lines[0][0], `the message [level severity="info"] {"FirstLevel":{"SecondLevel":"ceiling"}}`)
}

func TestLogWithStructuredData(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)
	data := syslog.StructuredData{
		"firstLevel": {"secondLevel": "the value"},
	}

	logger.LogWith(syslog.INFO, "the message", data)

	lines := openAndReadFile(testFile)

	assert.Contains(t, lines[0][0], `the message [firstLevel secondLevel="the value"][level severity="info"]`)
}

func TestLogEmergency(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Emergency("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="emerg"\] $`, lines[0][0])
}

func TestLogEmergencyWithData(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.EmergencyWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="emerg"\] {"key":"value"}$`, lines[0][0])
}

func TestLogAlert(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Alert("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="alert"\] $`, lines[0][0])
}

func TestLogAlertWithData(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.AlertWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="alert"\] {"key":"value"}$`, lines[0][0])
}

func TestLogCritical(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Critical("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="crit"\] $`, lines[0][0])
}

func TestLogCriticalWithData(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.CriticalWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="crit"\] {"key":"value"}$`, lines[0][0])
}

func TestLogError(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Error("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="err"\] $`, lines[0][0])
}

func TestLogErrorWithData(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.ErrorWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="err"\] {"key":"value"}$`, lines[0][0])
}

func TestLogWarning(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Warning("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="warning"\] $`, lines[0][0])
}

func TestLogWarningWithData(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.WarningWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="warning"\] {"key":"value"}$`, lines[0][0])
}

func TestLogNotice(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Notice("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="notice"\] $`, lines[0][0])
}

func TestLogNoticeWithData(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.NoticeWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="notice"\] {"key":"value"}$`, lines[0][0])
}

func TestLogInfo(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Info("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="info"\] $`, lines[0][0])
}

func TestLogInfoWithData(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.InfoWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="info"\] {"key":"value"}$`, lines[0][0])
}

func TestLogDebug(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Debug("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="debug"\] $`, lines[0][0])
}

func TestLogDebugWithData(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.DebugWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `the message \[level severity="debug"\] {"key":"value"}$`, lines[0][0])
}

func TestLogWithMinLevel(t *testing.T) {
	setUp()
	logger := loggers.Syslog{Path: testFile, MinLevel: syslog.INFO}

	logger.Debug("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 0)
}

func TestLogSameLevelAsMinLevel(t *testing.T) {
	setUp()
	logger := getLogger(testFile, 1)

	logger.Info("the message")

	lines := openAndReadFile(testFile)
	assert.Len(t, lines, 1)
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
	defer func() { file.Close() }()
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
