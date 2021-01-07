package http

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	net "net/http"
)

var NoRequestBodyDecoderFoundError = errors.New("unsupported content type").
	Status(net.StatusUnsupportedMediaType).
	Level(log_level.DEBUG)
