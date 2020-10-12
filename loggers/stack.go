package loggers

import (
	"errors"
	"github.com/lanvard/contract/inter"
)

type Stack struct {
	Loggers []string
	app     inter.Maker
}

func (s Stack) SetApp(app inter.Maker) inter.Logger {
	s.app = app
	return s
}

func (s Stack) Log(severity inter.Severity, message string) {
	for _, logger := range s.getLoggers() {
		logger.Log(severity, message)
	}
}

func (s Stack) LogWith(severity inter.Severity, message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.LogWith(severity, message, context)
	}
}

func (s Stack) Emergency(message string) {
	for _, logger := range s.getLoggers() {
		logger.Emergency(message)
	}
}

func (s Stack) EmergencyWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.EmergencyWith(message, context)
	}
}

func (s Stack) Alert(message string) {
	for _, logger := range s.getLoggers() {
		logger.Alert(message)
	}
}

func (s Stack) AlertWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.AlertWith(message, context)
	}
}

func (s Stack) Critical(message string) {
	for _, logger := range s.getLoggers() {
		logger.Critical(message)
	}
}

func (s Stack) CriticalWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.CriticalWith(message, context)
	}
}

func (s Stack) Error(message string) {
	for _, logger := range s.getLoggers() {
		logger.Error(message)
	}
}

func (s Stack) ErrorWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.ErrorWith(message, context)
	}
}

func (s Stack) Warning(message string) {
	for _, logger := range s.getLoggers() {
		logger.Warning(message)
	}
}

func (s Stack) WarningWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.WarningWith(message, context)
	}
}

func (s Stack) Notice(message string) {
	for _, logger := range s.getLoggers() {
		logger.Notice(message)
	}
}

func (s Stack) NoticeWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.NoticeWith(message, context)
	}
}

func (s Stack) Info(message string) {
	for _, logger := range s.getLoggers() {
		logger.Info(message)
	}
}

func (s Stack) InfoWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.InfoWith(message, context)
	}
}

func (s Stack) Debug(message string) {
	for _, logger := range s.getLoggers() {
		logger.Debug(message)
	}
}

func (s Stack) DebugWith(message string, context interface{}) {
	for _, logger := range s.getLoggers() {
		logger.DebugWith(message, context)
	}
}

func (s Stack) getLoggers() []inter.Logger {
	var loggers []inter.Logger
	allLoggers := s.app.Make("config.Logging.Loggers").(map[string]interface{})

	for _, loggerName := range s.Loggers {
		logger, ok := allLoggers[loggerName]
		if !ok {
			panic(errors.New("no logger found by: " + loggerName))
		}
		loggers = append(loggers, logger.(inter.Logger))
	}

	return loggers
}
