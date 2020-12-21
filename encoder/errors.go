package encoder

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	"net/http"
)

var EncodeError = errors.New("").Status(http.StatusInternalServerError).Level(log_level.EMERGENCY)
