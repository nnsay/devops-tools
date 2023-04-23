package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type SlackBlockText struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type SlackBlock struct {
	Type string         `json:"type"`
	Text SlackBlockText `json:"text"`
}

type SlackMessage struct {
	Channel string       `json:"channel"`
	Text    string       `json:"text,omitempty"`
	Blocks  []SlackBlock `json:"blocks"`
}

func SendNotification(channel string, title string, message string) {
	slackHook, _ := os.LookupEnv("SLACK_HOOK")
	fmt.Printf("Stack hook url: %s\n", slackHook)

	// https://api.slack.com/reference/surfaces/formatting
	slackMessage := SlackMessage{
		Channel: channel, Blocks: []SlackBlock{
			{
				Type: "section",
				Text: SlackBlockText{
					Type: "mrkdwn",
					// https://github.com/iamcal/emoji-data
					Text: fmt.Sprintf(":sos: *%s*", title),
				},
			},
			{
				Type: "section",
				Text: SlackBlockText{
					Type: "mrkdwn",
					Text: message,
				},
			},
		},
	}
	messageData, _ := json.Marshal(slackMessage)
	// https://api.slack.com/messaging/webhooks
	request, _ := http.NewRequest(http.MethodPost, slackHook, bytes.NewBuffer(messageData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	client.Do(request)
}
