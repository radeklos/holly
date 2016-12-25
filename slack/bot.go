package slack

import "net/http"

type SlackBot struct {
}

func (s *SlackBot) Do(chat ChatInterface) {
	http.Get(chat.Build())
}
