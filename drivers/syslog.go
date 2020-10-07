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
	case string:
		rawData = data
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

// Log that the system is unusable
func (r Syslog) Emergency(message string) {
	r.Log(syslog.EMERG, message)
}

// Log that the system is unusable
func (r Syslog) EmergencyWith(message string, data interface{}) {
	r.LogWith(syslog.EMERG, message, data)
}

// A condition that should be corrected immediately, such as a corrupted system database.
func (r Syslog) Alert(message string) {
	r.Log(syslog.ALERT, message)
}

// A condition that should be corrected immediately, such as a corrupted system database. w
func (r Syslog) AlertWith(message string, data interface{}) {
	r.LogWith(syslog.ALERT, message, data)
}

// Critical conditions
func (r Syslog) Critical(message string) {
	r.Log(syslog.CRIT, message)
}

// Critical conditions
func (r Syslog) CriticalWith(message string, data interface{}) {
	r.LogWith(syslog.CRIT, message, data)
}

// Error conditions
func (r Syslog) Error(message string) {
	r.Log(syslog.ERR, message)
}

// Error conditions
func (r Syslog) ErrorWith(message string, data interface{}) {
	r.LogWith(syslog.ERR, message, data)
}

// Warning conditions
func (r Syslog) Warning(message string) {
	r.Log(syslog.WARNING, message)
}

// Warning conditions
func (r Syslog) WarningWith(message string, data interface{}) {
	r.LogWith(syslog.WARNING, message, data)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (r Syslog) Notice(message string) {
	r.Log(syslog.NOTICE, message)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (r Syslog) NoticeWith(message string, data interface{}) {
	r.LogWith(syslog.NOTICE, message, data)
}

// Informational messages
func (r Syslog) Info(message string) {
	r.Log(syslog.INFO, message)
}

// Informational messages
func (r Syslog) InfoWith(message string, data interface{}) {
	r.LogWith(syslog.INFO, message, data)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (r Syslog) Debug(message string) {
	r.Log(syslog.DEBUG, message)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (r Syslog) DebugWith(message string, data interface{}) {
	r.LogWith(syslog.DEBUG, message, data)
}
