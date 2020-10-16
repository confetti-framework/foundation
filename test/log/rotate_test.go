package log

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/foundation"
	"github.com/lanvard/foundation/loggers"
	"github.com/lanvard/syslog"
	"github.com/stretchr/testify/assert"
	"github.com/vigneshuvi/GoDateFormat"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNameWithDateSoItCanRotate(t *testing.T) {
	setUp()
	var logger inter.Logger = loggers.Syslog{Path: testDir + "{yyyy-mm-dd}_test.log", MinLevel: syslog.INFO}
	logger = logger.SetApp(newTestApp())
	logger.Info("the message")

	dateWithCorrectFormat := time.Now().Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	assert.FileExists(t, testDir+dateWithCorrectFormat+"_test.log")
}

func TestDonNotRemoveLatestFile(t *testing.T) {
	// Given
	setUp()
	var logger inter.Logger = loggers.Syslog{Path: testDir + "{yyyy-mm-dd}_test.log", MinLevel: syslog.INFO}
	logger = logger.SetApp(newTestApp())
	logger.Info("the message")

	// When
	logger.Clear()

	// Then
	assert.Len(t, getFiles(), 1)
}

func TestDontRemoveOtherLogs(t *testing.T) {
	// Given
	setUp()
	logger1 := getLogger(testDir+"1_test.log", 2)
	logger1.Info("old message")

	logger2 := getLogger(testDir+"2_test.log", 2)
	logger2.Info("new message")

	logger3 := getLogger(testDir+"3_test.log", 2)
	logger3.Info("new message")

	// When
	logger1.Clear()

	// Then
	files := getFiles()
	assert.Len(t, files, 3)
	assert.Equal(t, testDir+"1_test.log", files[0])
	assert.Equal(t, testDir+"2_test.log", files[1])
	assert.Equal(t, testDir+"3_test.log", files[2])
}

func TestRemoveSecondFileIfMaxOne(t *testing.T) {
	// Given
	setUp()
	aLogFileIsPresent(testDir + "2019-10-21_test.log")
	logger := getLogger(testDir+"{yyyy-mm-dd}_test.log", 1)
	logger.Info("new message")

	// When
	logger.Clear()

	// Then
	assert.Len(t, getFiles(), 1)
}

func TestOnlyRemoveCurrentChannelFiles(t *testing.T) {
	// Given
	setUp()
	aLogFileIsPresent(testDir + "2019-10-21_channel1.log")
	aLogFileIsPresent(testDir + "2019-10-21_channel2.log")
	logger := getLogger(testDir+"{yyyy-mm-dd}_channel2.log", 1)
	logger.Info("new message")

	// When
	logger.Clear()

	// Then
	assert.Len(t, getFiles(), 2)
}

func aLogFileIsPresent(filename string) {
	logger := getLogger(filename, 1)
	logger.Info("message in " + filename)
}

func getFiles() []string {
	var files []string

	err := filepath.Walk(testDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func getLogger(file string, maxFiles int) inter.Logger {
	var logger inter.Logger = loggers.Syslog{Path: file, MinLevel: syslog.DEBUG, MaxFiles: maxFiles}
	logger = logger.SetApp(newTestApp())
	return logger
}

func getLoggerWithType(file string, messageType string) inter.Logger {
	var logger inter.Logger = loggers.Syslog{Path: file, MinLevel: syslog.DEBUG, Type: messageType}
	logger = logger.SetApp(newTestApp())
	return logger
}

func newTestApp() *foundation.Application {
	app := foundation.NewApp()
	app.Bind("config.App.Name", "testApp")
	return app
}
