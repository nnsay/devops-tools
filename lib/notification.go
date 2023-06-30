package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type SlackText struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji bool   `json:"emoji,omitempty"`
}

type SlackBlock struct {
	Type   string       `json:"type"`
	Text   *SlackText   `json:"text,omitempty"`
	Fields *[]SlackText `json:"fields,omitempty"`
}

type SlackMessage struct {
	Channel string        `json:"channel"`
	Text    string        `json:"text,omitempty"`
	Blocks  *[]SlackBlock `json:"blocks,omitempty"`
}

func doSendNotice(slackHook string, messageData []byte) {
	// https://api.slack.com/messaging/webhooks
	request, _ := http.NewRequest(http.MethodPost, slackHook, bytes.NewBuffer(messageData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	client.Do(request)
}

func SendNotification(channel string, title string, messages []SlackBlock) {
	slackHook, _ := os.LookupEnv("SLACK_HOOK")
	fmt.Printf("Stack hook url: %s\n", slackHook)

	// https://api.slack.com/reference/surfaces/formatting
	blocks := []SlackBlock{
		{
			Type: "header",
			Text: &SlackText{
				Type: "plain_text",
				Text: title,
				// https://github.com/iamcal/emoji-data
				Emoji: true,
			},
		},
	}
	blocks = append(blocks, messages...)
	slackMessage := SlackMessage{
		Channel: channel,
		Blocks:  &blocks,
	}
	messageData, _ := json.Marshal(slackMessage)
	fmt.Println("slack data: ", string(messageData))
	doSendNotice(slackHook, messageData)
}
