package view_helper

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	"net/http"
)

var NoSuchTemplateError = errors.
	New("customize providers.ViewServiceProvider to load the correct templates").
	Status(http.StatusInternalServerError).
	Level(log_level.EMERGENCY)
