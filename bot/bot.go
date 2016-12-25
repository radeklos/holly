package bot

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorhill/cronexpr"

	"github.com/radeklos/holly/slack"
)

type CronMessage struct {
	CronLine string
	Channel  string
	Message  string
}

type Config struct {
	SlackToken   string
	CronMessages []CronMessage
}

func NewBot(c Config) *Bot {
	log.Printf("Creating bot instance")
	return &Bot{
		config:   c,
		slackBot: slack.SlackBot{},
	}
}

type Bot struct {
	config   Config
	slackBot slack.SlackBot
}

func (d *Bot) sendMessage(channel string, msg string) {
	d.slackBot.Do(
		slack.ChatPostMessage{
			Chat:    &slack.Chat{Token: d.config.SlackToken},
			Channel: channel,
			Text:    msg,
		},
	)
}

func (d *Bot) registerCronMessage(message CronMessage) {
	defer func() {
		if err := recover(); err != nil { //catch
			log.Error("message cannot be scheduled: ", err)
		}
	}()

	scheduleAction("message to "+message.Channel, cronexpr.MustParse(message.CronLine), func() {
		d.sendMessage(message.Channel, message.Message)
	})

	log.WithFields(log.Fields{
		"at":      message.CronLine,
		"text":    message.Message,
		"channel": message.Channel,
	}).Infof("message scheduled")
}

func (d *Bot) Run() {
	log.Println("Running bot...")
	for _, message := range d.config.CronMessages {
		d.registerCronMessage(message)
	}
}

func scheduleAction(name string, cron *cronexpr.Expression, f func()) {
	go func() {
		// Create a ticker function
		ticker := func() *time.Ticker {
			next := cron.Next(time.Now())
			diff := next.Sub(time.Now())
			log.Infof("next execution for %s will be at %s", name, next.Format(time.RFC1123))
			return time.NewTicker(diff)
		}
		// Run the Ticker
		tkr := ticker()
		for {
			<-tkr.C
			f()
			tkr = ticker()
		}
	}()
}
