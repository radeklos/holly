package slack

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func NewBot(token string) *SlackBot {
	return &SlackBot{
		Token: token,
	}
}

type Message struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
	AsUser  bool   `json:"as_user"`
	Parse   string `json:"parse"`
}

type SlackBot struct {
	Token string
}

type SlackResponse struct {
	Ok    bool   `json:"ok"`
	Stuff string `json:"stuff,omitempty"`
	Error string `json:"error,omitempty"`
}

func (s *SlackBot) PostMessage(msg Message) {
	msg.AsUser = true
	msg.Parse = "full"
	url := "https://slack.com/api/chat.postMessage"

	b, err := json.Marshal(msg)
	if err != nil {
		log.Error("cannot marshal request body:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Authorization", strings.Join([]string{"Bearer", s.Token}, " "))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var response SlackResponse
	if err := json.Unmarshal(body, &response); err != nil {
		log.Error("cannot unmarshal response body:", err)
		return
	}
	if !response.Ok {
		log.Infof("error response:", string(body))
	}
}
