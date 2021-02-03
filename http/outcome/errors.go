package outcome

import (
	"github.com/confetti-framework/errors"
	net "net/http"
)

var FileNotFoundError = errors.New("file not found").Status(net.StatusNotFound)
var CanNotDownloadDirectoryError = FileNotFoundError
