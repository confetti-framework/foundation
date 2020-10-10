package log

import (
	"github.com/lanvard/foundation/loggers"
	"github.com/lanvard/syslog"
	"github.com/stretchr/testify/assert"
	"github.com/vigneshuvi/GoDateFormat"
	"testing"
	"time"
)

func TestNameWithDateSoItCanRotate(t *testing.T) {
	setUp()
	logger := loggers.Syslog{Testing: t, Path: testDir + "{yyyy-mm-dd}-log_test.log", MinLevel: syslog.INFO}

	logger.Info("the message")

	dateWithCorrectFormat := time.Now().Format(GoDateFormat.ConvertFormat("yyyy-mm-dd"))
	assert.FileExists(t, testDir+dateWithCorrectFormat+"-log_test.log")
}
