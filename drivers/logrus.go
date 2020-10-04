package drivers

import (
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

func (r *LogRus) init() {
	if r.Permissions == 0 {
		r.Permissions = 0644
	}
}

func (r LogRus) Log(level logrus.Level, message interface{}) {
	r.init()

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

	logger.WithFields(logrus.Fields{"test": "bla"}).Log(level, message)
}
