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
	var logger inter.Logger = loggers.Syslog{Path: testDir + "{yyyy-mm-dd}_log_test.log", MinLevel: syslog.INFO}
	logger = logger.SetApp(foundation.NewApp())
	logger.Info("the message")

	dateWithCorrectFormat := time.Now().Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	assert.FileExists(t, testDir+dateWithCorrectFormat+"_log_test.log")
}

func TestDonNotRemoveLatestFile(t *testing.T) {
	// Given
	setUp()
	var logger inter.Logger = loggers.Syslog{Path: testDir + "{yyyy-mm-dd}_log_test.log", MinLevel: syslog.INFO}
	logger = logger.SetApp(foundation.NewApp())
	logger.Info("the message")

	// When
	logger.Clear()

	// Then
	assert.Len(t, getFiles(), 1)
}

func TestRemoveSecondFileIfMaxOne(t *testing.T) {
	// Given
	setUp()
	logger := getLogger("1_log_test.log", 1)
	logger.Info("old message")
	logger2 := getLogger("2_log_test.log", 1)
	logger2.Info("old message")
	logger2.Info("new message")

	// When
	logger.Clear()

	// Then
	assert.Len(t, getFiles(), 1)
}

func TestRemoveWithMaxThree(t *testing.T) {
	// Given
	setUp()
	logger1 := getLogger("1_log_test.log", 3)
	logger1.Info("old message")

	logger2 := getLogger("2_log_test.log", 3)
	logger2.Info("new message")

	logger3 := getLogger("3_log_test.log", 3)
	logger3.Info("new message")

	logger4 := getLogger("4_log_test.log", 3)
	logger4.Info("new message")

	logger5 := getLogger("5_log_test.log", 3)
	logger5.Info("new message")

	// When
	logger1.Clear()

	// Then
	files := getFiles()
	assert.Len(t, files, 5)
	assert.Equal(t, testDir+"1_log_test.log", files[0])
	assert.Equal(t, testDir+"2_log_test.log", files[1])
	assert.Equal(t, testDir+"3_log_test.log", files[2])
	assert.Equal(t, testDir+"4_log_test.log", files[2])
	assert.Equal(t, testDir+"5_log_test.log", files[2])
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

func getLogger(path string, maxFiles int) inter.Logger {
	var logger inter.Logger = loggers.Syslog{Path: testDir + path, MinLevel: syslog.INFO, MaxFiles: maxFiles}
	app := foundation.NewApp()
	app.Bind("config.App.Name", "testApp")
	logger = logger.SetApp(app)
	return logger
}
