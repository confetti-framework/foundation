package loggers

import (
	"encoding/json"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/syslog"
	"github.com/vigneshuvi/GoDateFormat"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

type Syslog struct {
	Path       string
	Facility   inter.Facility
	Writer     io.Writer
	Permission os.FileMode
	MinLevel   inter.Severity
	MaxFiles   int
	app        inter.Maker
}

func (r Syslog) Clear() {
	// No files will be deleted when MaxFiles is 0
	if r.MaxFiles == 0 {
		return
	}

	files := getFilesByPath(r.Path)

	for i, file := range files {
		if i >= r.MaxFiles {
			err := os.Remove(filepath.Dir(r.Path) + string(os.PathSeparator) + file.Name())
			if err != nil {
				panic(err)
			}
		}
	}
}

func (r Syslog) SetApp(app inter.Maker) inter.Logger {
	r.app = app
	return r
}

func (r Syslog) Log(severity inter.Severity, message string) {
	r.LogWith(severity, message, "")
}

func (r Syslog) LogWith(severity inter.Severity, message string, context interface{}) {
	if r.MinLevel < severity {
		return
	}

	structuredData := syslog.StructuredData{}
	var rawData string

	switch context := context.(type) {
	case syslog.StructuredData:
		structuredData = context
	case string:
		rawData = context
	default:
		rawDataBytes, _ := json.Marshal(context)
		rawData = string(rawDataBytes)
	}

	structuredData["level"] = syslog.SDElement{"severity": syslog.KeyBySeverity(severity)}

	r.init().Log(
		severity,
		message,
		structuredData,
		rawData,
	)
}

// Log that the system is unusable
func (r Syslog) Emergency(message string) {
	r.Log(syslog.EMERG, message)
}

// Log that the system is unusable
func (r Syslog) EmergencyWith(message string, context interface{}) {
	r.LogWith(syslog.EMERG, message, context)
}

// A condition that should be corrected immediately, such as a corrupted system contextbase.
func (r Syslog) Alert(message string) {
	r.Log(syslog.ALERT, message)
}

// A condition that should be corrected immediately, such as a corrupted system contextbase. w
func (r Syslog) AlertWith(message string, context interface{}) {
	r.LogWith(syslog.ALERT, message, context)
}

// Critical conditions
func (r Syslog) Critical(message string) {
	r.Log(syslog.CRIT, message)
}

// Critical conditions
func (r Syslog) CriticalWith(message string, context interface{}) {
	r.LogWith(syslog.CRIT, message, context)
}

// Error conditions
func (r Syslog) Error(message string) {
	r.Log(syslog.ERR, message)
}

// Error conditions
func (r Syslog) ErrorWith(message string, context interface{}) {
	r.LogWith(syslog.ERR, message, context)
}

// Warning conditions
func (r Syslog) Warning(message string) {
	r.Log(syslog.WARNING, message)
}

// Warning conditions
func (r Syslog) WarningWith(message string, context interface{}) {
	r.LogWith(syslog.WARNING, message, context)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (r Syslog) Notice(message string) {
	r.Log(syslog.NOTICE, message)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (r Syslog) NoticeWith(message string, context interface{}) {
	r.LogWith(syslog.NOTICE, message, context)
}

// Informational messages
func (r Syslog) Info(message string) {
	r.Log(syslog.INFO, message)
}

// Informational messages
func (r Syslog) InfoWith(message string, context interface{}) {
	r.LogWith(syslog.INFO, message, context)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (r Syslog) Debug(message string) {
	r.Log(syslog.DEBUG, message)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (r Syslog) DebugWith(message string, context interface{}) {
	r.LogWith(syslog.DEBUG, message, context)
}

func (r Syslog) init() syslog.Logger {
	hostname, _ := os.Hostname()
	if r.Writer == nil {
		r.Writer = fileWriter(r)
	}

	appName := r.app.Make("config.App.Name").(string)
	procid := strconv.Itoa(os.Getpid())

	return syslog.NewLogger(r.Writer, r.Facility, hostname, appName, procid)
}

func fileWriter(r Syslog) *os.File {
	if r.Permission == 0 {
		r.Permission = 0644
	}

	// We overwrite the default value of 0.
	if r.Facility == 0 {
		r.Facility = syslog.USER
	}

	// create extra dir if needed
	err := os.MkdirAll(filepath.Dir(r.Path), r.Permission)
	if err != nil {
		panic(err)
	}
	fileName := getDynamicFileName(r.Path)

	writer, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, r.Permission)
	if err != nil {
		panic(err)
	}
	return writer
}

func getDynamicFileName(rawFileName string) string {
	r := regexp.MustCompile(`(?P<prefix>.*)(?P<braces_r>{)(?P<date_format>.*?)(?P<braces_l>})(?P<suffix>.*)`)
	match := r.FindStringSubmatch(rawFileName)

	if len(match) < 5 {
		return rawFileName
	}

	result := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	return result["prefix"] + time.Now().Format(GoDateFormat.ConvertFormat(result["date_format"])) + result["suffix"]
}

func getFilesByPath(filePath string) []os.FileInfo {
	var files []os.FileInfo

	dir, fileName := filepath.Split(filePath)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// logic to check info.IsDir(), info.Name(), and written actual filename
		if info.IsDir() {
			return nil
		}

		re := regexp.MustCompile(`(\{.*\})`)
		regexFilename := re.ReplaceAllString(fileName, ".*")

		r := regexp.MustCompile(regexFilename)
		if !r.Match([]byte(info.Name())) {
			return nil
		}
		files = append(files, info)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
