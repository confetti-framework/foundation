package loggers

import (
	"github.com/lanvard/contract/inter"
	"github.com/sirupsen/logrus"
)

type Slack struct {
	Url      string
	Username string
	Emoji    string
	Level    logrus.Level
}

func (s Slack) Log(severity inter.Severity, message string) {
	panic("implement me")
}

func (s Slack) LogWith(severity inter.Severity, message string, data interface{}) {
	panic("implement me")
}

func (s Slack) Emergency(message string) {
	panic("implement me")
}

func (s Slack) EmergencyWith(message string, data interface{}) {
	panic("implement me")
}

func (s Slack) Alert(message string) {
	panic("implement me")
}

func (s Slack) AlertWith(message string, data interface{}) {
	panic("implement me")
}

func (s Slack) Critical(message string) {
	panic("implement me")
}

func (s Slack) CriticalWith(message string, data interface{}) {
	panic("implement me")
}

func (s Slack) Error(message string) {
	panic("implement me")
}

func (s Slack) ErrorWith(message string, data interface{}) {
	panic("implement me")
}

func (s Slack) Warning(message string) {
	panic("implement me")
}

func (s Slack) WarningWith(message string, data interface{}) {
	panic("implement me")
}

func (s Slack) Notice(message string) {
	panic("implement me")
}

func (s Slack) NoticeWith(message string, data interface{}) {
	panic("implement me")
}

func (s Slack) Info(message string) {
	panic("implement me")
}

func (s Slack) InfoWith(message string, data interface{}) {
	panic("implement me")
}

func (s Slack) Debug(message string) {
	panic("implement me")
}

func (s Slack) DebugWith(message string, data interface{}) {
	panic("implement me")
}
