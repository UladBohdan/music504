package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type lastfmTrack struct {
	Artist    map[string]string `json:"artist"`
	Name      string            `json:"name"`
	Playcount string            `json:"playcount"`
	URL       string            `json:"url"`
}

type lastfmChart map[string]map[string][]lastfmTrack

func getLastfmChart(username string, dateFrom, dateTo time.Time) (string, error) {
	tsFrom := strconv.FormatInt(dateFrom.Unix(), 10)
	tsTo := strconv.FormatInt(dateTo.Unix(), 10)
	url := fmt.Sprintf(lastfmAPIUrl, username, tsFrom, tsTo)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.String(), nil
}

func getTopTrack(username string, from, to time.Time) (track, error) {
	str, err := getLastfmChart(username, from, to)
	if err != nil {
		return track{}, err
	}

	var chart lastfmChart
	json.Unmarshal([]byte(str), &chart)

	top := chart["weeklytrackchart"]["track"][0]

	playcount, err := strconv.ParseInt(top.Playcount, 10, 64)
	if err != nil {
		return track{}, err
	}

	return track{
		Artist:    top.Artist["#text"],
		Name:      top.Name,
		Playcount: int(playcount),
		URL:       top.URL,
		username:  username,
	}, nil
}
