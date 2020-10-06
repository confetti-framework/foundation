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

var mock = struct{ FirstLevel interface{} }{
	struct{ SecondLevel string }{
		"ceiling",
	},
}

func TestCreateLogFileIfNotExists(t *testing.T) {
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Log(syslog.INFO, "some content")

	assert.FileExists(t, testFile)
}

func TestGetAddOneErrorLine(t *testing.T) {
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Log(syslog.INFO, "error line 1")

	assert.Len(t, openAndReadFile(testFile), 1)
}

func TestGetAddTwoErrorLines(t *testing.T) {
	logger := drivers.Syslog{Testing: t, Path: testFile}

	logger.Log(syslog.INFO, "error line 1")
	logger.Log(syslog.INFO, "error line 2")

	assert.Len(t, openAndReadFile(testFile), 2)
}

// func TestContentOfError(t *testing.T) {
// 	logger := drivers.Syslog{Testing: t, Path: testFile}
//
// 	logger.Log(syslog.INFO, "error line 1")
// 	logger.Log(syslog.INFO, "error line 2")
//
// 	lines := openAndReadFile(testFile)
// 	assert.Contains(t, lines[0][0], "time=")
// 	assert.Contains(t, lines[0][0], "msg=\"error line 1\"")
// 	assert.Contains(t, lines[1][0], "time=")
// 	assert.Contains(t, lines[1][0], "msg=\"error line 2\"")
// }
//
// func TestLogWithLevel(t *testing.T) {
// 	logger := drivers.Syslog{Testing: t, Path: testFile}
//
// 	logger.Log(syslog.INFO, "info message")
//
// 	lines := openAndReadFile(testFile)
// 	assert.Contains(t, lines[0][0], "level=info msg=\"info message\"")
// }
//
// func TestLogWithoutFormatter(t *testing.T) {
// 	logger := drivers.Syslog{Testing: t, Path: testFile}
//
// 	logger.Log(syslog.INFO, "info message")
//
// 	lines := openAndReadFile(testFile)
// 	assert.Regexp(t, "^time=\"[^\"]*\" level=info msg=\"info message\"$", lines[0][0])
// }
//
// func TestLogWithFormatterWithoutQuotes(t *testing.T) {
// 	logger := drivers.Syslog{
// 		Testing:   t,
// 		Path:      testFile,
// 		Formatter: &syslog.TextFormatter{DisableQuote: true},
// 	}
//
// 	logger.Log(syslog.INFO, "info message")
//
// 	lines := openAndReadFile(testFile)
// 	assert.Regexp(t, "level=info msg=info message$", lines[0][0])
// }
//
// func TestLogWithString(t *testing.T) {
// 	logger := drivers.Syslog{Testing: t, Path: testFile}
//
// 	logger.LogWith(syslog.INFO, "the info", "string data")
//
// 	lines := openAndReadFile(testFile)
// 	assert.Contains(t, lines[0][0], `data="\"string data\""`)
// }
//
// func TestLogWithMap(t *testing.T) {
// 	logger := drivers.Syslog{Testing: t, Path: testFile}
//
// 	logger.LogWith(syslog.INFO, "the info", map[string]string{"key": "value"})
//
// 	lines := openAndReadFile(testFile)
// 	assert.Contains(t, lines[0][0], `data="{\"key\":\"value\"}"`)
// }
//
// func TestLogWithStruct(t *testing.T) {
// 	logger := drivers.Syslog{Testing: t, Path: testFile}
//
// 	logger.LogWith(syslog.INFO, "the info", mock)
//
// 	lines := openAndReadFile(testFile)
// 	assert.Contains(t, lines[0][0], `data="{\"FirstLevel\":{\"SecondLevel\":\"ceiling\"}}"`)
// }
//
// func TestLogLevels(t *testing.T) {
// 	logger := drivers.Syslog{Testing: t, Path: testFile}
//
// 	logger.LogWith(syslog.INFO, "the info", mock)
//
// 	lines := openAndReadFile(testFile)
// 	assert.Contains(t, lines[0][0], `data="{\"FirstLevel\":{\"SecondLevel\":\"ceiling\"}}"`)
// }
//
// func TestTrace(t *testing.T) {
// 	logger := drivers.Syslog{Testing: t, Path: testFile}
//
// 	logger.Trace("the info")
//
// 	lines := openAndReadFile(testFile)
// 	assert.Contains(t, lines[0][0], `data="{\"FirstLevel\":{\"SecondLevel\":\"ceiling\"}}"`)
// }

func openAndReadFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Failed to open file: %s", fileName)
	}
	lines, err := readFile(file)
	if err != nil {
		fmt.Printf("Failed to read file: %s", fileName)
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
