package drivers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

type LogRus struct {
	Path        string
	Permissions os.FileMode
	Level       logrus.Level
	Days        int
	Testing     *testing.T
	Formatter   logrus.Formatter
}

func (r LogRus) init() *logrus.Logger {
	if r.Permissions == 0 {
		r.Permissions = 0644
	}

	// Delete the file later if a test is used
	if r.Testing != nil {
		r.Testing.Cleanup(func() {
			_ = os.Remove(r.Path)
		})
	}

	logger := logrus.New()
	if r.Formatter != nil {
		logger.Formatter = r.Formatter
	}

	file, err := os.OpenFile(r.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, r.Permissions)
	if err == nil {
		logger.Out = file
	}

	return logger
}

func (r LogRus) Log(level logrus.Level, message string) {
	r.init().Log(level, message)
}

func (r LogRus) LogWith(level logrus.Level, message string, data interface{}) {
	result, _ := json.Marshal(data)
	r.init().WithField("data", string(result)).Log(level, message)
}
