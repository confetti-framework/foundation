package report

import (
	"github.com/lanvard/errors"
	"github.com/lanvard/syslog/log_level"
	net "net/http"
)

var EncodeError = errors.New("").Status(net.StatusInternalServerError).Level(log_level.EMERGENCY)
var NoSuchTemplate = EncodeError.Wrap("customize providers.ViewServiceProvider to load more templates")
