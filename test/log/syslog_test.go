package log

import (
	"encoding/csv"
	"fmt"
	"github.com/lanvard/foundation/drivers"
	"github.com/lanvard/syslog"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"os"
	"testing"
)

const testDir = "/tmp/"
const testFile = testDir + "log_test.log"

var structMock = struct{ FirstLevel interface{} }{
	struct{ SecondLevel string }{
		"ceiling",
	},
}

func TestCreateLogFileIfNotExists(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Log(syslog.INFO, "some content")

	assert.FileExists(t, testFile)
}

func TestGetAddOneErrorLine(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Log(syslog.INFO, "error line 1")

	assert.Len(t, openAndReadFile(testFile), 1)
}

func TestGetAddTwoErrorLines(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Log(syslog.INFO, "error line 1")
	logger.Log(syslog.INFO, "error line 2")

	assert.Len(t, openAndReadFile(testFile), 2)
}

func TestLogWithLevel(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Log(syslog.INFO, "message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, "^<14.*info:", lines[0][0])
}

func TestContentOfError(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Log(syslog.INFO, "error line 1")
	logger.Log(syslog.INFO, "error line 2")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, `^<14>1 \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.* - info: error line 1 -`, lines[0][0])
	assert.Regexp(t, `^<14>1 \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.* - info: error line 2 -`, lines[1][0])
}

func TestLogWithString(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.LogWith(syslog.INFO, "the message", "string data")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - info: the message - string data$`, lines[0][0])
}

func TestLogWithMap(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.LogWith(syslog.INFO, "the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Contains(t, lines[0][0], ` - - info: the message - {"key":"value"}`)
}

func TestLogWithStruct(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.LogWith(syslog.INFO, "the message", structMock)

	lines := openAndReadFile(testFile)
	assert.Contains(t, lines[0][0], ` - - info: the message - {"FirstLevel":{"SecondLevel":"ceiling"}}`)
}

func TestLogLevels(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.LogWith(syslog.INFO, "the message", structMock)

	lines := openAndReadFile(testFile)
	assert.Contains(t, lines[0][0], ` - - info: the message - {"FirstLevel":{"SecondLevel":"ceiling"}}`)
}

func TestLogWithStructuredData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}
	data := syslog.StructuredData{
		"firstLevel": {"secondLevel": "the value"},
	}

	logger.LogWith(syslog.INFO, "the message", data)

	lines := openAndReadFile(testFile)

	assert.Contains(t, lines[0][0], ` - - info: the message [firstLevel secondLevel="the value"]`)
}

func TestLogEmergency(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Emergency("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - emerg: the message - $`, lines[0][0])
}

func TestLogEmergencyWithData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.EmergencyWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - emerg: the message - {"key":"value"}$`, lines[0][0])
}

func TestLogAlert(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Alert("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - alert: the message - $`, lines[0][0])
}

func TestLogAlertWithData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.AlertWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - alert: the message - {"key":"value"}$`, lines[0][0])
}

func TestLogCritical(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Critical("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - crit: the message - $`, lines[0][0])
}

func TestLogCriticalWithData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.CriticalWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - crit: the message - {"key":"value"}$`, lines[0][0])
}

func TestLogError(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Error("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - err: the message - $`, lines[0][0])
}

func TestLogErrorWithData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.ErrorWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - err: the message - {"key":"value"}$`, lines[0][0])
}

func TestLogWarning(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Warning("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - warning: the message - $`, lines[0][0])
}

func TestLogWarningWithData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.WarningWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - warning: the message - {"key":"value"}$`, lines[0][0])
}

func TestLogNotice(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Notice("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - notice: the message - $`, lines[0][0])
}

func TestLogNoticeWithData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.NoticeWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - notice: the message - {"key":"value"}$`, lines[0][0])
}

func TestLogInfo(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Info("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - info: the message - $`, lines[0][0])
}

func TestLogInfoWithData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.InfoWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - info: the message - {"key":"value"}$`, lines[0][0])
}

func TestLogDebug(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Debug("the message")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - debug: the message - $`, lines[0][0])
}

func TestLogDebugWithData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.DebugWith("the message", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - debug: the message - {"key":"value"}$`, lines[0][0])
}

func setUp() {
	_ = os.Remove(testFile)
}

func openAndReadFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %s\n", fileName)
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
