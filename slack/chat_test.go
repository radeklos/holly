package slack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChatPostMessage(t *testing.T) {
	f := ChatPostMessage{
		Chat:    &Chat{Token: "token"},
		Channel: "channel",
		Text:    "hello world",
	}

	assert.Equal(t, "https://slack.com/api/chat.postMessage?as_user=true&channel=channel&parse=full&text=hello+world&token=token", f.Build())
}
