package main

import (
	"flag"
	"fmt"
	"sort"
)

var (
	enableTips = flag.Bool("enable_tips", false, "Enable tips on the bottom of each message")
	prodPtr    = flag.Bool("prod", false, "Enables pushing to prod channel. Be careful!")
	timePeriod = flag.String("time_period", "twoWeeks", "Defines the time period to grab stats for")

	debug bool // equals !prod
)

type track struct {
	Artist    string
	Name      string
	Playcount int
	URL       string

	username string
}

func run() {
	period := parseTimePeriod(*timePeriod)
	from, to := timeframeToTime(period)

	topTracks := make([]track, len(lastfmUsernames))
	for i, username := range lastfmUsernames {
		tr, err := getTopTrack(username, from, to)
		if err != nil {
			fmt.Printf("failed to get top track for '%s': %s\n", username, err)
			return
		}
		topTracks[i] = tr
	}
	sort.Slice(topTracks, func(i, j int) bool {
		return topTracks[i].Playcount > topTracks[j].Playcount
	})

	message := formatMessage(topTracks, period, from, to)

	err := publishTracks(message)
	if err != nil {
		fmt.Printf("failed to push to telegram: %s\n", err)
		return
	}
}

func main() {
	flag.Parse()
	debug = !*prodPtr
	run()
}
