package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PlexItem struct {
	Title string `xml:"title,attr"`
	Type  string `xml:"type,attr"`
}

type PlexService struct {
	url       string
	apiKey    string
	Watchlist []PlexItem
}

func (plex *PlexService) WaitUntilAvailable() {
	waitUntilServiceAvailable(plex.url)
}

func (plex *PlexService) fetchPlexWatchlist() {
	url := fmt.Sprintf("%s/library/sections/watchlist/all?X-Plex-Token=%s", plex.url, plex.apiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch Plex watchlist: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read Plex watchlist response: %v", err)
	}

	var result struct {
		Directory []PlexItem `xml:"Directory"`
		Video     []PlexItem `xml:"Video"`
	}
	err = xml.Unmarshal(body, &result)
	if err != nil {
		log.Fatalf("Failed to unmarshal Plex watchlist XML: %v", err)
	}

	log.Printf("Found %d directories and %d videos in Plex watchlist", len(result.Directory), len(result.Video))

	plex.Watchlist = append(result.Directory, result.Video...)
}

func (plex *PlexService) GetMovies() []PlexItem {
	if len(plex.Watchlist) == 0 {
		plex.fetchPlexWatchlist()
	}

	fitlerFunc := func(item PlexItem) bool { return item.Type == "movie" }
	return filter(plex.Watchlist, fitlerFunc)
}

func (plex *PlexService) GetShows() []PlexItem {
	if len(plex.Watchlist) == 0 {
		plex.fetchPlexWatchlist()
	}
	fitlerFunc := func(item PlexItem) bool { return item.Type == "show" }
	return filter(plex.Watchlist, fitlerFunc)
}
