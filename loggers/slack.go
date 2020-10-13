package loggers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/syslog"
	"net/http"
	"time"
)

type Slack struct {
	WebhookUrl string
	MinLevel   inter.Severity
}

func (s Slack) SetApp(_ inter.Maker) inter.Logger {
	return s
}

func (s Slack) Log(severity inter.Severity, message string) {
	s.LogWith(severity, message, "")
}

func (s Slack) LogWith(severity inter.Severity, message string, data interface{}) {
	if s.MinLevel < severity {
		return
	}

	if s.WebhookUrl == "" {
		panic(errors.New("no URL found for Slack logger"))
	}

	sendSlackNotification(s.WebhookUrl, syslog.KeyBySeverity(severity)+": "+message, data)
}

// Log that the system is unusable
func (s Slack) Emergency(message string) {
	s.Log(syslog.EMERG, message)
}

// Log that the system is unusable
func (s Slack) EmergencyWith(message string, context interface{}) {
	s.LogWith(syslog.EMERG, message, context)
}

// A condition that should be corrected immediately, such as a corrupted system contextbase.
func (s Slack) Alert(message string) {
	s.Log(syslog.ALERT, message)
}

// A condition that should be corrected immediately, such as a corrupted system contextbase. w
func (s Slack) AlertWith(message string, context interface{}) {
	s.LogWith(syslog.ALERT, message, context)
}

// Critical conditions
func (s Slack) Critical(message string) {
	s.Log(syslog.CRIT, message)
}

// Critical conditions
func (s Slack) CriticalWith(message string, context interface{}) {
	s.LogWith(syslog.CRIT, message, context)
}

// Error conditions
func (s Slack) Error(message string) {
	s.Log(syslog.ERR, message)
}

// Error conditions
func (s Slack) ErrorWith(message string, context interface{}) {
	s.LogWith(syslog.ERR, message, context)
}

// Warning conditions
func (s Slack) Warning(message string) {
	s.Log(syslog.WARNING, message)
}

// Warning conditions
func (s Slack) WarningWith(message string, context interface{}) {
	s.LogWith(syslog.WARNING, message, context)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (s Slack) Notice(message string) {
	s.Log(syslog.NOTICE, message)
}

// Normal but significant conditions
// Conditions that are not error conditions, but that may require special handling.
func (s Slack) NoticeWith(message string, context interface{}) {
	s.LogWith(syslog.NOTICE, message, context)
}

// Informational messages
func (s Slack) Info(message string) {
	s.Log(syslog.INFO, message)
}

// Informational messages
func (s Slack) InfoWith(message string, context interface{}) {
	s.LogWith(syslog.INFO, message, context)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (s Slack) Debug(message string) {
	s.Log(syslog.DEBUG, message)
}

// Debug-level messages
// Messages containing information that is normally only useful when debugging a program.
func (s Slack) DebugWith(message string, context interface{}) {
	s.LogWith(syslog.DEBUG, message, context)
}

type SlackRequestBody struct {
	Text string `json:"text"`
}

// @todo use notifications https://github.com/lanvard/lanvard/issues/70
//
// sendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and extra data.
func sendSlackNotification(webhookUrl string, msg string, rawData interface{}) {
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
