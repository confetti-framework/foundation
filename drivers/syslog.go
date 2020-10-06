package drivers

import (
	"github.com/lanvard/syslog"
	"os"
	"testing"
)

type Severity = syslog.Priority
type Facility = syslog.Priority

type Syslog struct {
	Path     string
	FileMode os.FileMode
	MinLevel Severity
	Days     int
	Testing  *testing.T
	Facility Facility
	AppName  string
	Procid   string
}

func (r Syslog) init() syslog.Logger {
	if r.FileMode == 0 {
		r.FileMode = 0644
	}

	// Delete the file later if a test is used
	if r.Testing != nil {
		r.Testing.Cleanup(func() {
			_ = os.Remove(r.Path)
		})
	}

	// We overwrite the default value of 0.
	if r.Facility == 0 {
		r.Facility = syslog.USER
	}
	hostname, _ := os.Hostname()
	file, err := os.OpenFile(r.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, r.FileMode)
	if err != nil {
		panic(err)
	}
	return syslog.NewLogger(file, r.Facility, hostname, r.AppName, "")
}

func (r Syslog) Log(severity Severity, message string) {
	r.init().Log(
		severity,
		syslog.KeyBySeverity(severity)+" "+message,
		nil,
		"",
	)
}
