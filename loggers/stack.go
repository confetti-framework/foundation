package loggers

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/syslog/level"
)

type Stack struct {
	Channels []string
	app      inter.Maker
}

func (s Stack) Clear() {}

func (s Stack) SetApp(app inter.Maker) inter.Logger {
	s.app = app
	return s
}

func (s Stack) Log(severity level.Level, message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Log(severity, message, arguments...)
	}
}

func (s Stack) LogWith(severity level.Level, message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.LogWith(severity, message, context)
	}
}

func (s Stack) Emergency(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Emergency(message, arguments...)
	}
}

func (s Stack) EmergencyWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.EmergencyWith(message, context)
	}
}

func (s Stack) Alert(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Alert(message, arguments...)
	}
}

func (s Stack) AlertWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.AlertWith(message, context)
	}
}

func (s Stack) Critical(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Critical(message, arguments...)
	}
}

func (s Stack) CriticalWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.CriticalWith(message, context)
	}
}

func (s Stack) Error(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Error(message, arguments...)
	}
}

func (s Stack) ErrorWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.ErrorWith(message, context)
	}
}

func (s Stack) Warning(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Warning(message, arguments...)
	}
}

func (s Stack) WarningWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.WarningWith(message, context)
	}
}

func (s Stack) Notice(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Notice(message, arguments...)
	}
}

func (s Stack) NoticeWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.NoticeWith(message, context)
	}
}

func (s Stack) Info(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Info(message, arguments...)
	}
}

func (s Stack) InfoWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.InfoWith(message, context)
	}
}

func (s Stack) Debug(message string, arguments ...interface{}) {
	for _, logger := range s.getLoggers() {
		logger.Debug(message, arguments...)
	}
}

func (s Stack) DebugWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.DebugWith(message, context)
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
