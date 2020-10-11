package loggers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/syslog"
	"github.com/sirupsen/logrus"
)

type Slack struct {
	Url      string
	Username string
	Emoji    string
	Level    logrus.Level
}

func (s Slack) Log(severity inter.Severity, message string) {
	s.LogWith(severity, message, "")
}

func (s Slack) LogWith(severity inter.Severity, message string, data interface{}) {
	panic("implement me")
}

// Log that the system is unusable
func (s Slack) Emergency(message string) {
	s.Log(syslog.EMERG, message)
}

// Log that the system is unusable
func (s Slack) EmergencyWith(message string, context interface{}) {
	s.LogWith(syslog.EMERG, message, context)
}

// A condition that should be corrected immediately, such as a corrupted system contextbase.
func (s Slack) Alert(message string) {
	s.Log(syslog.ALERT, message)
}

// A condition that should be corrected immediately, such as a corrupted system contextbase. w
func (s Slack) AlertWith(message string, context interface{}) {
	s.LogWith(syslog.ALERT, message, context)
}

// Critical conditions
func (s Slack) Critical(message string) {
	s.Log(syslog.CRIT, message)
}

// Critical conditions
func (s Slack) CriticalWith(message string, context interface{}) {
	s.LogWith(syslog.CRIT, message, context)
}

// Error conditions
func (s Slack) Error(message string) {
	s.Log(syslog.ERR, message)
}

// Error conditions
func (s Slack) ErrorWith(message string, context interface{}) {
	s.LogWith(syslog.ERR, message, context)
}

// Warning conditions
func (s Slack) Warning(message string) {
	s.Log(syslog.WARNING, message)
}

// Warning conditions
func (s Slack) WarningWith(message string, context interface{}) {
	s.LogWith(syslog.WARNING, message, context)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (s Slack) Notice(message string) {
	s.Log(syslog.NOTICE, message)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (s Slack) NoticeWith(message string, context interface{}) {
	s.LogWith(syslog.NOTICE, message, context)
}

// Informational messages
func (s Slack) Info(message string) {
	s.Log(syslog.INFO, message)
}

// Informational messages
func (s Slack) InfoWith(message string, context interface{}) {
	s.LogWith(syslog.INFO, message, context)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (s Slack) Debug(message string) {
	s.Log(syslog.DEBUG, message)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (s Slack) DebugWith(message string, context interface{}) {
	s.LogWith(syslog.DEBUG, message, context)
}
