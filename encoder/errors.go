package encoder

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	"net/http"
)

var EncodeError = errors.WithLevel(errors.WithStatus(
	errors.New(""),
	http.StatusInternalServerError),
	log_level.EMERGENCY)
