package main

import (
	"fmt"
	"net/url"
)

func produceYandexMusicLink(artist, name string) string {
	s := url.QueryEscape(artist + " " + name)
	return fmt.Sprintf(yandexMusicSearchURL, s)
}
