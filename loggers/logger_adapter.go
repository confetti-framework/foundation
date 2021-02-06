package loggers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/syslog/log_level"
)

type loggerFacade struct {
	Logger inter.Logger
}

func NewLoggerFacade(logger inter.Logger) *loggerFacade {
	return &loggerFacade{Logger: logger}
}

func (l *loggerFacade) SetApp(app inter.AppReader) inter.LoggerFacade {
	l.Logger = l.Logger.SetApp(app)
	return l
}

func (l *loggerFacade) Group(group string) inter.LoggerFacade {
	l.Logger = l.Logger.Group(group)
	return l
}

func (l *loggerFacade) Log(severity log_level.Level, message string, arguments ...interface{}) {
	l.Logger.Log(severity, message, arguments...)
}

func (l *loggerFacade) LogWith(severity log_level.Level, message string, data interface{}) {
	l.Logger.LogWith(severity, message, data)
}

// Emergency Log that the system is unusable
func (l *loggerFacade) Emergency(message string, arguments ...interface{}) {
	l.Logger.Log(log_level.EMERGENCY, message, arguments...)
}

// EmergencyWith Log that the system is unusable
func (l *loggerFacade) EmergencyWith(message string, data interface{}) {
	l.Logger.LogWith(log_level.EMERGENCY, message, data)
}

// Alert A condition that should be corrected immediately.
func (l *loggerFacade) Alert(message string, arguments ...interface{}) {
	l.Logger.Log(log_level.ALERT, message, arguments...)
}

// AlertWith A condition that should be corrected immediately.
func (l *loggerFacade) AlertWith(message string, data interface{}) {
	l.Logger.LogWith(log_level.ALERT, message, data)
}

// Critical conditions
func (l *loggerFacade) Critical(message string, arguments ...interface{}) {
	l.Logger.Log(log_level.CRITICAL, message, arguments...)
}

// CriticalWith Critical conditions
func (l *loggerFacade) CriticalWith(message string, data interface{}) {
	l.Logger.LogWith(log_level.CRITICAL, message, data)
}

// Error conditions
func (l *loggerFacade) Error(message string, arguments ...interface{}) {
	l.Logger.Log(log_level.ERROR, message, arguments...)
}

// ErrorWith Error conditions
func (l *loggerFacade) ErrorWith(message string, data interface{}) {
	l.Logger.LogWith(log_level.ERROR, message, data)
}

// Warning conditions
func (l *loggerFacade) Warning(message string, arguments ...interface{}) {
	l.Logger.Log(log_level.WARNING, message, arguments...)
}

// WarningWith Warning conditions
func (l *loggerFacade) WarningWith(message string, data interface{}) {
	l.Logger.LogWith(log_level.WARNING, message, data)
}

// Notice Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (l *loggerFacade) Notice(message string, arguments ...interface{}) {
	l.Logger.Log(log_level.NOTICE, message, arguments...)
}

// NoticeWith Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (l *loggerFacade) NoticeWith(message string, data interface{}) {
	l.Logger.LogWith(log_level.NOTICE, message, data)
}

// Info Informational messages
func (l *loggerFacade) Info(message string, arguments ...interface{}) {
	l.Logger.Log(log_level.INFO, message, arguments...)
}

// InfoWith Informational messages
func (l *loggerFacade) InfoWith(message string, data interface{}) {
	l.Logger.LogWith(log_level.INFO, message, data)
}

// Debug Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (l *loggerFacade) Debug(message string, arguments ...interface{}) {
	l.Logger.Log(log_level.DEBUG, message, arguments...)
}

// DebugWith Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (l *loggerFacade) DebugWith(message string, data interface{}) {
	l.Logger.LogWith(log_level.DEBUG, message, data)
}
