package drivers

import "github.com/sirupsen/logrus"

type Slack struct {
	Url      string
	Username string
	Emoji    string
	Level    logrus.Level
}
