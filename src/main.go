package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/robfig/cron/v3"
)

type DbResult struct {
	ID int `json:"id"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

var client = &http.Client{
	Timeout: 5 * time.Second, // Set a timeout of 5 seconds
}

var (
	LANGUAGE_PROFILE = 1
)

func main() {
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("Config: %+v\n", config)

	plex := PlexService{
		url:    config.Plex.URL,
		apiKey: config.Plex.ApiKey,
	}

	plex.WaitUntilAvailable()

	config.Radarr.ProcessWatchList(plex.GetMovies())

	config.Sonarr.ProcessWatchList(plex.GetShows())

	c := cron.New()

	c.AddFunc("*/5 * * * *", func() {
		config.Radarr.ProcessWatchList(plex.GetMovies())

		config.Sonarr.ProcessWatchList(plex.GetShows())
	})
	c.Start()

	select {}
}
