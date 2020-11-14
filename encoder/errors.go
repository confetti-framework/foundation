package encoder

import (
	"github.com/lanvard/errors"
	"github.com/lanvard/syslog/log_level"
	"net/http"
)

var EncodeError = errors.New("").Status(http.StatusInternalServerError).Level(log_level.EMERGENCY)
