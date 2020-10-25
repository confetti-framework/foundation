package loggers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/errors"
	"github.com/lanvard/syslog"
	"github.com/lanvard/syslog/log_level"
	"net/http"
	"time"
)

type Slack struct {
	WebhookUrl string
	MinLevel   log_level.Level
}

func (s Slack) SetApp(_ inter.Maker) inter.Logger {
	return s
}

func (s Slack) Clear() {}

func (s Slack) Log(severity log_level.Level, message string, arguments ...interface{}) {
	s.LogWith(severity, fmt.Sprintf(message, arguments...), "")
}

func (s Slack) LogWith(severity log_level.Level, message string, data interface{}) {
	if s.MinLevel < severity {
		return
	}

	if s.WebhookUrl == "" {
		panic(errors.New("no URL found for Slack logger"))
	}

	sendSlackNotification(s.WebhookUrl, syslog.KeyBySeverity(severity)+": "+message, data)
}

// Log that the system is unusable
func (s Slack) Emergency(message string, arguments ...interface{}) {
	s.Log(log_level.EMERGENCY, message, arguments...)
}

// Log that the system is unusable
func (s Slack) EmergencyWith(message string, context interface{}) {
	s.LogWith(log_level.EMERGENCY, message, context)
}

// A condition that should be corrected immediately, such as a corrupted system contextbase.
func (s Slack) Alert(message string, arguments ...interface{}) {
	s.Log(log_level.ALERT, message, arguments...)
}

// A condition that should be corrected immediately, such as a corrupted system contextbase. w
func (s Slack) AlertWith(message string, context interface{}) {
	s.LogWith(log_level.ALERT, message, context)
}

// Critical conditions
func (s Slack) Critical(message string, arguments ...interface{}) {
	s.Log(log_level.CRITICAL, message, arguments...)
}

// Critical conditions
func (s Slack) CriticalWith(message string, context interface{}) {
	s.LogWith(log_level.CRITICAL, message, context)
}

// Error conditions
func (s Slack) Error(message string, arguments ...interface{}) {
	s.Log(log_level.ERROR, message, arguments...)
}

// Error conditions
func (s Slack) ErrorWith(message string, context interface{}) {
	s.LogWith(log_level.ERROR, message, context)
}

// Warning conditions
func (s Slack) Warning(message string, arguments ...interface{}) {
	s.Log(log_level.WARNING, message, arguments...)
}

// Warning conditions
func (s Slack) WarningWith(message string, context interface{}) {
	s.LogWith(log_level.WARNING, message, context)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (s Slack) Notice(message string, arguments ...interface{}) {
	s.Log(log_level.NOTICE, message, arguments...)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (s Slack) NoticeWith(message string, context interface{}) {
	s.LogWith(log_level.NOTICE, message, context)
}

// Informational messages
func (s Slack) Info(message string, arguments ...interface{}) {
	s.Log(log_level.INFO, message, arguments...)
}

// Informational messages
func (s Slack) InfoWith(message string, context interface{}) {
	s.LogWith(log_level.INFO, message, context)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (s Slack) Debug(message string, arguments ...interface{}) {
	s.Log(log_level.DEBUG, message, arguments...)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (s Slack) DebugWith(message string, context interface{}) {
	s.LogWith(log_level.DEBUG, message, context)
}

type SlackRequestBody struct {
	Text string `json:"text"`
}

// @todo use notifications https://github.com/lanvard/lanvard/issues/70
//
// sendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and extra data.
func sendSlackNotification(webhookUrl string, msg string, rawData interface{}) {
	// Don't send stack trace
	if _, ok := rawData.(error); ok {
		return
	}

	if rawData == nil {
		data, _ := json.Marshal(rawData)
		msg = msg + "\n" + string(data)
	}
	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}
	if buf.String() != "ok" {
		panic(errors.New("non-ok response returned from Slack"))
	}

	return
}
