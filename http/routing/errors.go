package routing

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	net "net/http"
)

var RouteError = errors.New("").Level(log_level.DEBUG)
var MethodNotAllowedError = RouteError.Wrap("HTTP method not allowed").Status(net.StatusMethodNotAllowed)
var RouteNotFoundError = RouteError.Wrap("no match was found for the specified URL").Status(net.StatusNotFound)
var AppNotFoundError = RouteError.Wrap("inter.App not found in RouteCollection").Status(net.StatusInternalServerError).Level(log_level.CRITICAL)
