package drivers

import (
	"github.com/sirupsen/logrus"
)

type LogRus struct {
	Path  string
	Level logrus.Level
	Days  int
}
