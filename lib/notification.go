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

type SlackTextBlock struct {
	Type string    `json:"type"`
	Text SlackText `json:"text,omitempty"`
}

type SlackFieldBlock struct {
	Type   string      `json:"type"`
	Fields []SlackText `json:"fields,omitempty"`
}

type SlackMessage struct {
	Channel string        `json:"channel"`
	Text    string        `json:"text,omitempty"`
	Blocks  []interface{} `json:"blocks"`
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

func SendNotification(channel string, title string, messages []interface{}) {
	slackHook, _ := os.LookupEnv("SLACK_HOOK")
	fmt.Printf("Stack hook url: %s\n", slackHook)

	// https://api.slack.com/reference/surfaces/formatting
	slackMessage := SlackMessage{
		Channel: channel,
		Blocks: []interface{}{
			SlackTextBlock{
				Type: "header",
				Text: SlackText{
					Type: "plain_text",
					// https://github.com/iamcal/emoji-data
					Text:  title,
					Emoji: true,
				},
			},
		},
	}
	slackMessage.Blocks = append(slackMessage.Blocks, messages...)
	messageData, _ := json.Marshal(slackMessage)
	// fmt.Print(string(messageData))
	doSendNotice(slackHook, messageData)
}
