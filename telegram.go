package main

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/telegram-bot-api.v4"
)

func formatMessage(tracks []track, period timeframe, from, to time.Time) string {
	var s string

	if debug {
		s += fmt.Sprint("/// SENT IN DEBUG MODE ///\n\n")
	}

	s += fmt.Sprintf("*%s*\n", periodToTitle[period])
	s += fmt.Sprintf("%v - %v\n\n", from.Format(outputTimeFormat), to.Format(outputTimeFormat))

	for _, tr := range tracks {
		s += fmt.Sprintf("[🎵](%v) %v :  [%v](%v)\n", produceYandexMusicLink(tr.Artist, tr.Name), tr.Artist, tr.Name, tr.URL)
		fmt.Println("XXX: ", produceYandexMusicLink(tr.Artist, tr.Name))
	}

	s += "\nlistened accordingly by:\n\n"

	for i, tr := range tracks {
		s += fmt.Sprintf("%v. *%v* - %v %v\n", i+1, names[tr.username], tr.Playcount, func() string {
			if i == 0 {
				return "times"
			}
			return ""
		}())
	}

	s += fmt.Sprintf("\n[👾](%v) ваш бот", "https://github.com/uladbohdan/music504")

	if *enableTips {
		s += fmt.Sprintf("\ntip: emojis are clickable!")
	}

	return s
}

func publishTracks(message string) error {
	bot, err := tgbotapi.NewBotAPI(telegramBotToken)
	if err != nil {
		return err
	}

	if debug {
		bot.Debug = true
		log.Printf("Authorized on account %s", bot.Self.UserName)
	}

	var channelName string
	if debug {
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
