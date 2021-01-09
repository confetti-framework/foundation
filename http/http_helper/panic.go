package http_helper

import "github.com/confetti-framework/errors"

func GetErrorFromPanic(recoverRaw interface{}) error {
	var err error
	switch rec := recoverRaw.(type) {
	case string:
		err = errors.New(rec)
	case error:
		err = rec
	default:
		err = errors.New("can't convert panic to response. Error or string required")
	}
	return err
}
