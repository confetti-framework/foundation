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

func (s Stack) getLoggers() []inter.Logger {
	var loggers []inter.Logger
	var err error
	allLoggers := s.app.Make("config.Logging.Channels").(map[string]interface{})

	for _, loggerName := range s.Channels {
		rawLogger, ok := allLoggers[loggerName]
		if !ok {
			err = errors.New("channel '%s' does not exist", loggerName)
			break
		} else {
			logger := rawLogger.(inter.Logger)
			logger = logger.SetApp(s.app)
			loggers = append(loggers, logger)
		}
	}

	if err != nil {
		panic(errors.WithMessage(err, s.prefixMessage()))
	}

	return loggers
}

func (s Stack) prefixMessage() string {
	if len(s.Channels) == 1 {
		return "can not log to channel"
	} else {
		return "can not log to one of the channels"
	}
}
