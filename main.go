package main

import (
	"flag"
	"fmt"
)

var (
	debug      = flag.Bool("debug", false, "Enables debug mode")
	timePeriod = flag.String("time_period", "twoWeeks", "Defines the time period to grab stats for")
)

type track struct {
	Artist    string
	Name      string
	Playcount string
	URL       string
}

func run() {
	topTracks := make(map[string]track)

	period := parseTimePeriod(*timePeriod)

	from, to := timeframeToTime(period)

	for _, username := range lastfmUsernames {
		tr, err := getTopTrack(username, from, to)
		if err != nil {
			fmt.Printf("failed to get top track for '%s': %s\n", username, err)
			return
		}
		topTracks[username] = tr
	}

	message := formatMessage(topTracks, period, from, to)

	err := publishTracks(message)
	if err != nil {
		fmt.Printf("failed to push to telegram: %s\n", err)
		return
	}
}

func main() {
	flag.Parse()
	run()
}
