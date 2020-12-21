package http

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	net "net/http"
)

var NoRequestBodyDecoderFoundError = errors.New("no request body decoder found. Check the headers and http method").
	Status(net.StatusUnsupportedMediaType).
	Level(log_level.DEBUG)
