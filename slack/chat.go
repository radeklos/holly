package slack

import (
	"net/url"
)

type ChatInterface interface {
	Build() string
}

type Chat struct {
	Token string
}

func (c Chat) Build() string {
	return "https://slack.com/api/chat."
}

type ChatPostMessage struct {
	*Chat
	Channel string
	Text    string
}

func (c ChatPostMessage) Build() string {
	v := url.Values{}
	v.Set("channel", c.Channel)
	v.Add("text", c.Text)
	v.Add("token", c.Chat.Token)
	v.Add("as_user", "true")
	v.Add("parse", "full")
	v.Encode()

	return c.Chat.Build() + "postMessage?" + v.Encode()
}
