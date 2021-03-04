package loggers

import (
	"encoding/json"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog"
	"github.com/confetti-framework/syslog/log_level"
	"github.com/vigneshuvi/GoDateFormat"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"
)

type Syslog struct {
	Path           string
	Facility       syslog.Facility
	group          string // MSGID intended for filtering
	Writer         io.Writer
	Permission     os.FileMode
	MinLevel       log_level.Level
	MaxFiles       int
	app            inter.AppReader
	HideStackTrace bool
}

func (r Syslog) SetApp(app inter.AppReader) inter.Logger {
	r.app = app
	return r
}

func (r Syslog) Group(group string) inter.Logger {
	r.group = group
	return r
}

func (r Syslog) Clear() bool {
	// No files will be deleted when MaxFiles is 0
	if r.MaxFiles == 0 {
		return false
	}

	files := getFilesToCleanUp(r.Path)

	for i, file := range files {
		if i >= r.MaxFiles {
			err := os.Remove(filepath.Dir(r.Path) + string(os.PathSeparator) + file.Name())
			if err != nil {
				panic(err)
			}
		}
	}

	return true
}

func (r Syslog) Log(severity log_level.Level, message string, arguments ...interface{}) {
	r.LogWith(severity, fmt.Sprintf(message, arguments...), "")
}

func (r Syslog) LogWith(severity log_level.Level, message string, rawContext interface{}) {
	if r.MinLevel < severity {
		return
	}

	structuredData := syslog.StructuredData{}
	var rawData string

	switch context := rawContext.(type) {
	case syslog.StructuredData:
		structuredData = context
	case error:
		if r.HideStackTrace {
			break
		}
		stack, ok := errors.FindStack(context)
		if !ok {
			break
		}
		rawData = fmt.Sprintf("%+v", stack)
	case string:
		rawData = context
	default:
		rawDataBytes, _ := json.Marshal(context)
		rawData = string(rawDataBytes)
	}

	structuredData["level"] = syslog.SDElement{"severity": syslog.KeyBySeverity(severity)}

	r.init().Log(
		severity,
		r.group,
		structuredData,
		message+" %s",
		rawData,
	)
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

func getFilesToCleanUp(filePath string) []os.FileInfo {
	var files []os.FileInfo

	dir, fileName := filepath.Split(filePath)

	// if no dynamic filename is present, then we don't have to clean up anything
	regexDynamic := regexp.MustCompile(`(\{.*\})`)
	if !regexDynamic.Match([]byte(fileName)) {
		return nil
	}

	// Determine a regex that tells us which files to clean up
	dynamicPart := regexDynamic.FindStringSubmatch(fileName)
	lengthOfDynamic := strconv.Itoa(len(dynamicPart[0]) - 2)
	regexCleanUp := regexDynamic.ReplaceAllString(fileName, `.{`+lengthOfDynamic+`}`)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// Ignore directories
		if info.IsDir() {
			return nil
		}

		//  If the file does not match, we should not clean it up
		r := regexp.MustCompile(regexCleanUp)
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
