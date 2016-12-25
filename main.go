package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/radeklos/holly/bot"
	"github.com/spf13/viper"
)

func EnvToCronMessages(data []string) []bot.CronMessage {
	var s []bot.CronMessage

	for _, item := range data {
		splits := strings.Split(item, "=")
		key := splits[0]
		message := splits[1]
		if strings.HasPrefix(key, "MSG") {
			parts := strings.SplitN(key, "_", 4)
			if len(parts) < 4 {
				log.WithFields(log.Fields{
					"name":  key,
					"value": message,
				}).Warn("Cannot understand cron message")
				continue
			}
			s = append(s, bot.CronMessage{
				CronLine: fmt.Sprintf("%s %s * * MON-FRI *", parts[2], parts[1]),
				Channel:  parts[3],
				Message:  message,
			})
		}
	}
	return s
}

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("SLACK_TOKEN", "")

	bot := bot.NewBot(
		bot.Config{
			SlackToken:   viper.GetString("SLACK_TOKEN"),
			CronMessages: EnvToCronMessages(os.Environ()),
		},
	)
	bot.Run()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	// Run forever unless we get a signal
	for sig := range signals {
		log.Println(sig)
		os.Exit(1)
	}
}
