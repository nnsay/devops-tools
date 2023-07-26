package lib_test

import (
	"testing"

	"github.com/nnsay/devops-tools/lib"
)

func TestNotice(t *testing.T) {
	channel := "#test"
	title := "Test Message"
	messages := []lib.SlackBlock{
		{
			Type: "section",
			Text: &lib.SlackText{
				Type: "mrkdwn",
				Text: ":test_tube: This is a message for test",
			},
		},
	}
	code := lib.SendNotification(channel, title, messages)
	t.Log(code)
}
