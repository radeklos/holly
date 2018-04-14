package api

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/radeklos/holly/slack"
)

type Message struct {
	Text    string `json:"text"`
	Channel string `json:"channel"`
}

func New(slackBot *slack.SlackBot) *Api {
	return &Api{
		SlackBot: slackBot,
	}
}

type Api struct {
	SlackBot *slack.SlackBot
}

func (a *Api) SendMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	_ = json.NewDecoder(r.Body).Decode(&message)
	go a.SlackBot.Send(slack.Message{
		Channel: message.Channel,
		Text:    message.Text,
	})
}

func (s *Api) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/send", s.SendMessage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}
