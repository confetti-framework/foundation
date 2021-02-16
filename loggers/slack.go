package loggers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/confetti-framework/contract/inter"
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog"
	"github.com/confetti-framework/syslog/log_level"
	"net/http"
	"time"
)

type Slack struct {
	WebhookUrl string
	MinLevel   log_level.Level
}

func (s Slack) SetApp(_ inter.AppReader) inter.Logger {
	return s
}

func (s Slack) Group(_ string) inter.Logger {
	// For Slack we don't use the group for anything
	return s
}

func (s Slack) Clear() bool {
	return false
}

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

type SlackRequestBody struct {
	Text string `json:"text"`
}

// @todo use notifications https://github.com/confetti-framework/confetti-framework/issues/70
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
