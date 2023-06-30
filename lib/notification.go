package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

// https://api.slack.com/messaging/webhooks
// func doSendNotice(slackHook string, messageData []byte) {
// 	request, _ := http.NewRequest(http.MethodPost, slackHook, bytes.NewBuffer(messageData))
// 	request.Header.Set("Content-Type", "application/json")
// 	client := &http.Client{
// 		Timeout: time.Second * 60,
// 	}
// 	response, err := client.Do(request)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer response.Body.Close()
// 	b, _ := io.ReadAll(response.Body)
// 	responseBody := string(b)
// 	fmt.Println(responseBody)
// }

// https://api.slack.com/methods/chat.postMessage
func postMessage(messageData []byte) int {
	url := "https://slack.com/api/chat.postMessage"
	slackToken, hasToken := os.LookupEnv("SLACK_TOKEN")
	if !hasToken {
		fmt.Println("SLACK_TOKEN not set")
		return -1
	}
	request, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(messageData))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", slackToken))
	client := &http.Client{
		Timeout: time.Second * 60,
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return response.StatusCode
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(response.Body)
		responseBody := string(b)
		fmt.Println(responseBody)
	}
	return response.StatusCode
}

func SendNotification(channel string, title string, messages []SlackBlock) int {
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
	// doSendNotice(slackHook, messageData)
	return postMessage(messageData)
}
