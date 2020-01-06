package foundation

import (
	"upspin.io/errors"
	"upspin.io/upspin"
)

type ApplicationError struct {
	Path upspin.PathName
	Op  errors.Op
	Kind errors.Kind
	Err error
}

func (e ApplicationError) Error() string {
	return "application error: " + e.Err.Error()
}
