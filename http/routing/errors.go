package routing

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	net "net/http"
)

var RouteError = errors.WithLevel(errors.New(""), log_level.DEBUG)
var MethodNotAllowedError = errors.WithStatus(errors.Wrap(RouteError,
	"HTTP method not allowed"),
	net.StatusMethodNotAllowed)
var RouteNotFoundError = errors.WithStatus(errors.Wrap(RouteError,
	"no match was found for the specified URL"),
	net.StatusNotFound)
var AppNotFoundError = errors.WithLevel(errors.WithStatus(errors.Wrap(RouteError,
	"inter.App not found in RouteCollection"),
	net.StatusInternalServerError),
	log_level.CRITICAL)
