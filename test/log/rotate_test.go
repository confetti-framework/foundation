package log

import (
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
	logger := loggers.Syslog{Testing: t, Path: testDir + "{yyyy-mm-dd}_log_test.log", MinLevel: syslog.INFO}

	logger.Info("the message")

	dateWithCorrectFormat := time.Now().Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	assert.FileExists(t, testDir+dateWithCorrectFormat+"_log_test.log")
}

func TestDonNotRemoveLatestFile(t *testing.T) {
	// Given
	setUp()
	logger := loggers.Syslog{Testing: t, Path: testDir + "{yyyy-mm-dd}_log_test.log", MinLevel: syslog.INFO}
	logger.Info("the message")

	// When
	logger.Clear()

	// Then
	assert.Len(t, getFiles(), 1)
}

func TestRemoveSecondFileIfMaxOne(t *testing.T) {
	// Given
	setUp()
	logger := loggers.Syslog{Testing: t, Path: testDir + "1_log_test.log", MinLevel: syslog.INFO, MaxFiles: 1}
	logger.Info("old message")
	logger.Path = testDir + "2_log_test.log"
	logger.Info("new message")

	// When
	logger.Clear()

	// Then
	assert.Len(t, getFiles(), 1)
}

func TestRemoveWithMaxThree(t *testing.T) {
	// Given
	setUp()
	logger := loggers.Syslog{Testing: t, MinLevel: syslog.INFO, MaxFiles: 3}
	logger.Path = testDir + "1_log_test.log"
	logger.Info("old message")
	logger.Path = testDir + "2_log_test.log"
	logger.Info("new message")
	logger.Path = testDir + "3_log_test.log"
	logger.Info("new message")
	logger.Path = testDir + "4_log_test.log"
	logger.Info("new message")
	logger.Path = testDir + "5_log_test.log"
	logger.Info("new message")

	// When
	logger.Clear()

	// Then
	files := getFiles()
	assert.Len(t, files, 3)
	assert.Equal(t, testDir+"1_log_test.log", files[0])
	assert.Equal(t, testDir+"2_log_test.log", files[1])
	assert.Equal(t, testDir+"3_log_test.log", files[2])
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
