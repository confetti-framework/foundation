package loggers

import (
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
)

type Stack struct {
	Channels []string
	app      inter.AppReader
	group    string
}

func (s Stack) SetApp(app inter.AppReader) inter.Logger {
	s.app = app
	return s
}

func (s Stack) Clear() {}

func (s Stack) Group(group string) inter.Logger {
	s.group = group

	return s
}

func (s Stack) Log(severity log_level.Level, message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).Log(severity, message, arguments...)
	}
}

func (s Stack) LogWith(severity log_level.Level, message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).LogWith(severity, message, context)
	}
}

func (s Stack) Emergency(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).Emergency(message, arguments...)
	}
}

func (s Stack) EmergencyWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).EmergencyWith(message, context)
	}
}

func (s Stack) Alert(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).Alert(message, arguments...)
	}
}

func (s Stack) AlertWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).AlertWith(message, context)
	}
}

func (s Stack) Critical(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).Critical(message, arguments...)
	}
}

func (s Stack) CriticalWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).CriticalWith(message, context)
	}
}

func (s Stack) Error(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).Error(message, arguments...)
	}
}

func (s Stack) ErrorWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).ErrorWith(message, context)
	}
}

func (s Stack) Warning(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).Warning(message, arguments...)
	}
}

func (s Stack) WarningWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).WarningWith(message, context)
	}
}

func (s Stack) Notice(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).Notice(message, arguments...)
	}
}

func (s Stack) NoticeWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).NoticeWith(message, context)
	}
}

func (s Stack) Info(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).Info(message, arguments...)
	}
}

func (s Stack) InfoWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).InfoWith(message, context)
	}
}

func (s Stack) Debug(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).Debug(message, arguments...)
	}
}

func (s Stack) DebugWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Group(s.group).DebugWith(message, context)
	}
}

func (s Stack) getLoggers() []inter.Logger {
	var loggers []inter.Logger
	allLoggers := s.app.Make("config.Logging.Channels").(map[string]interface{})

	for _, loggerName := range s.Channels {
		rawLogger, ok := allLoggers[loggerName]
		if !ok {
			panic(errors.New("can not log to one of the channels. Channel '" + loggerName + "' does not exist"))
		}
		logger := rawLogger.(inter.Logger)
		logger = logger.SetApp(s.app)
		loggers = append(loggers, logger)
	}

	return loggers
}
