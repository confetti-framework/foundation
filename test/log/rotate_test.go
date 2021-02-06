package log

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/foundation"
	"github.com/confetti-framework/foundation/loggers"
	"github.com/confetti-framework/syslog/log_level"
	"github.com/stretchr/testify/require"
	"github.com/vigneshuvi/GoDateFormat"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func Test_name_with_date_so_it_can_rotate(t *testing.T) {
	setUp()
	var logger inter.Logger = loggers.Syslog{Path: testDir + "{yyyy-mm-dd}_test.log", MinLevel: log_level.INFO}
	logger = logger.SetApp(newTestApp())
	logger.Log(log_level.INFO, "the message")

	dateWithCorrectFormat := time.Now().Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	require.FileExists(t, testDir+dateWithCorrectFormat+"_test.log")
}

func Test_don_not_remove_latest_file(t *testing.T) {
	// Given
	setUp()
	var logger inter.Logger = loggers.Syslog{Path: testDir + "{yyyy-mm-dd}_test.log", MinLevel: log_level.INFO}
	logger = logger.SetApp(newTestApp())
	logger.Log(log_level.INFO, "the message")

	// When
	logger.Clear()

	// Then
	require.Len(t, getFiles(), 1)
}

func Test_dont_remove_other_logs(t *testing.T) {
	// Given
	setUp()
	logger1 := getLogger(testDir+"1_test.log", 2)
	logger1.Log(log_level.INFO, "old message")

	logger2 := getLogger(testDir+"2_test.log", 2)
	logger2.Log(log_level.INFO, "new message")

	logger3 := getLogger(testDir+"3_test.log", 2)
	logger3.Log(log_level.INFO, "new message")

	// When
	logger1.Clear()

	// Then
	files := getFiles()
	require.Len(t, files, 3)
	require.Equal(t, testDir+"1_test.log", files[0])
	require.Equal(t, testDir+"2_test.log", files[1])
	require.Equal(t, testDir+"3_test.log", files[2])
}

func Test_remove_second_file_if_max_one(t *testing.T) {
	// Given
	setUp()
	aLogFileIsPresent(testDir + "2019-10-21_test.log")
	logger := getLogger(testDir+"{yyyy-mm-dd}_test.log", 1)
	logger.Log(log_level.INFO, "new message")

	// When
	logger.Clear()

	// Then
	require.Len(t, getFiles(), 1)
}

func Test_only_remove_current_channel_files(t *testing.T) {
	// Given
	setUp()
	aLogFileIsPresent(testDir + "2019-10-21_channel1.log")
	aLogFileIsPresent(testDir + "2019-10-21_channel2.log")
	logger := getLogger(testDir+"{yyyy-mm-dd}_channel2.log", 1)
	logger.Log(log_level.INFO, "new message")

	// When
	logger.Clear()

	// Then
	require.Len(t, getFiles(), 2)
}

func aLogFileIsPresent(filename string) {
	logger := getLogger(filename, 1)
	logger.Log(log_level.INFO, "message in "+filename)
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
	var logger inter.Logger = loggers.Syslog{Path: file, MinLevel: log_level.DEBUG, MaxFiles: maxFiles}
	logger = logger.SetApp(newTestApp())
	return logger
}

func newTestApp() *foundation.Application {
	app := foundation.NewApp()
	app.Bind("config.App.Name", "testApp")
	return app
}
