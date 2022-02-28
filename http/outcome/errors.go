package outcome

import (
	"github.com/confetti-framework/errors"
	net "net/http"
)

var FileNotFoundError = errors.WithStatus(errors.New("file not found"), net.StatusNotFound)
