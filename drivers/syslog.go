package drivers

import (
	"encoding/json"
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
	r.LogWith(severity, message, "")
}

func (r Syslog) LogWith(severity Severity, message string, data interface{}) {
	var structuredData syslog.StructuredData
	var rawData string
	switch data := data.(type) {
	case syslog.StructuredData:
		structuredData = data
	default:
		rawDataBytes, _ := json.Marshal(data)
		rawData = string(rawDataBytes)
	}

	r.init().Log(
		severity,
		syslog.KeyBySeverity(severity)+": "+message,
		structuredData,
		rawData,
	)
}
