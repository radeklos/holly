package slack

import (
	"net/http"
)

func NewBot(token string) *SlackBot {
	return &SlackBot{
		Token: token,
	}
}

type Message struct {
	Text    string
	Channel string
}

type SlackBot struct {
	Token string
}

func (s *SlackBot) Do(chat ChatInterface) {
	http.Get(chat.Build())
}

func (s *SlackBot) Send(msg Message) {
	s.Do(
		ChatPostMessage{
			Chat:    &Chat{Token: s.Token},
			Channel: msg.Channel,
			Text:    msg.Text,
		},
	)
}
