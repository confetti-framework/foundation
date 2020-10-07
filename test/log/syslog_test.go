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

	logger.LogWith(syslog.INFO, "the info", "string data")

	lines := openAndReadFile(testFile)
	assert.Regexp(t, ` - - info: the info - "string data"`, lines[0][0])
}

func TestLogWithMap(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.LogWith(syslog.INFO, "the info", map[string]string{"key": "value"})

	lines := openAndReadFile(testFile)
	assert.Contains(t, lines[0][0], ` - - info: the info - {"key":"value"}`)
}

func TestLogWithStruct(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.LogWith(syslog.INFO, "the info", structMock)

	lines := openAndReadFile(testFile)
	assert.Contains(t, lines[0][0], ` - - info: the info - {"FirstLevel":{"SecondLevel":"ceiling"}}`)
}

func TestLogLevels(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.LogWith(syslog.INFO, "the info", structMock)

	lines := openAndReadFile(testFile)
	assert.Contains(t, lines[0][0], ` - - info: the info - {"FirstLevel":{"SecondLevel":"ceiling"}}`)
}

func TestLogWithStructuredData(t *testing.T) {
	setUp()
	logger := drivers.Syslog{Testing: t, Path: testFile}
	data := syslog.StructuredData{
		"firstLevel": {"secondLevel": "the value"},
	}

	logger.LogWith(syslog.INFO, "the info", data)

	lines := openAndReadFile(testFile)

	assert.Contains(t, lines[0][0], ` - - info: the info [firstLevel secondLevel="the value"]`)
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
