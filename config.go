package main

type timeframe int

// A list of supported time frames to get stats for.
const (
	oneWeek timeframe = iota
	twoWeeks
	oneMonth
	threeMonths
	sixMonths
	oneYear
)

const (
	lastfmAPIKey = "2c0e360975c35880df3dcf3d4e3f769e"

	// Last.fm API call: https://www.last.fm/api/show/user.getWeeklyTrackChart
	lastfmAPIUrl = "http://ws.audioscrobbler.com/2.0/?method=user.getweeklytrackchart&format=json&api_key=2c0e360975c35880df3dcf3d4e3f769e&user=%s&from=%s&to=%s"

	telegramBotToken = "600356809:AAGd2kArl_QaKWlowo7TXQnLO2pN70laLcs"

	channelDebugName    = "@uladbohdan_test0"
	channelOfficialName = "@room504music"

	outputTimeFormat = "Jan 02"

	yandexMusicSearchURL = "https://music.yandex.ru/search?text=%v"
)

var (
	lastfmUsernames = []string{"uladbohdan", "stanis1r", "andrejkaAZAZ"}

	names = map[string]string{
		"uladbohdan":   "Уладзік",
		"stanis1r":     "Стасік",
		"andrejkaAZAZ": "Андрэйка",
	}

	periodToTitle = map[timeframe]string{
		oneWeek:     "Падборка за апошні тыдзень:",
		twoWeeks:    "Пятнічная падборка за два тыдні:",
		oneMonth:    "Падборачка за месяц",
		threeMonths: "Падборачка за тры месяцы!",
		sixMonths:   "За паўгады:",
		oneYear:     "🥂 гадавая!",
	}
)
