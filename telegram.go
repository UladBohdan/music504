package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/telegram-bot-api.v4"
)

func formatMessage(tracks map[string]track, period timeframe, from, to time.Time) string {
	var s string

	s += fmt.Sprintf("*%s*\n", periodToTitle[period])
	s += fmt.Sprintf("%d.%d - %d.%d    // incl.-excl.\n\n", from.Day(), from.Month(), to.Day(), to.Month())

	for username, track := range tracks {
		s += fmt.Sprintf("*%s* - %s скроблаў\n", names[username], track.Playcount)
		s += fmt.Sprintf("%s :  [%s](%s)\n\n", track.Artist, track.Name, track.URL)
	}

	s += fmt.Sprintf("ваш бот [♥](%s)", "https://github.com/uladbohdan/music504")
	return s
}

func publishTracks(message string) error {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		return err
	}

	if *debug {
		bot.Debug = true
		log.Printf("Authorized on account %s", bot.Self.UserName)
	}

	var channelName string
	if *debug {
		channelName = channelDebugName
	} else {
		channelName = channelOfficialName
	}

	msg := tgbotapi.NewMessageToChannel(channelName, message)
	msg.ParseMode = "Markdown"
	msg.DisableWebPagePreview = true

	bot.Send(msg)

	return nil
}
