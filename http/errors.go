package http

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	net "net/http"
)

var NoRequestBodyDecoderFoundError = errors.WithStatus(errors.WithLevel(errors.New("unsupported content type or HTTP method"),
	log_level.DEBUG),
	net.StatusUnsupportedMediaType)
