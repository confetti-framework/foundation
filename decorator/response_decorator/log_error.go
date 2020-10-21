package response_decorator

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/syslog/level"
)

type LogError struct{}

func (l LogError) Decorate(response inter.Response) inter.Response {
	if err, ok := response.Content().(error); ok {
		var logLevel level.Level
		if err, ok := err.(interface{ GetLevel() level.Level }); ok {
			logLevel = err.GetLevel()
		}
		response.App().Log().LogWith(logLevel, err.Error(), err)
	}

	return response
}
