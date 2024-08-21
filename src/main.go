package main

import (
	"log"
	"net/http"

	// "os"
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
	// // Open the log file for writing
	// logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatalf("Failed to open log file: %v", err)
	// }
	// defer logFile.Close()

	// // Set log output to the file
	// log.SetOutput(logFile)

	// Load configuration
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	log.Printf("Config: %+v\n", config)

	plex := PlexService{
		url:    config.Plex.URL,
		apiKey: config.Plex.ApiKey,
	}

	plex.WaitUntilAvailable()

	config.Radarr.ProcessWatchList(plex.GetMovies())
	config.Sonarr.ProcessWatchList(plex.GetShows())

	c := cron.New()

	c.AddFunc("*/5 * * * *", func() {
		log.Println("Fetching Plex watchlist")
		plex.fetchPlexWatchlist()

		log.Println("Processing Radarr watchlist")
		config.Radarr.ProcessWatchList(plex.GetMovies())

		log.Println("Processing Sonarr watchlist")
		config.Sonarr.ProcessWatchList(plex.GetShows())
	})
	c.Start()

	log.Println("Service started")
	select {}
}
