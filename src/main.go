package main

import (
	"flag"
	"log"

	"github.com/robfig/cron/v3"
)

// Config struct (define your structure here as per the config file)
type Config struct {
	Overseer struct {
		APIKey  string
		BaseURL string
	}
	Radarr struct {
		APIKey  string
		BaseURL string
	}
	Sonarr struct {
		APIKey  string
		BaseURL string
	}
}

func main() {
	// Define flags for each configuration option
	overseerAPIKey := flag.String("overseer-api-key", "", "API key for Overseer")
	overseerBaseURL := flag.String("overseer-base-url", "", "Base URL for Overseer")
	radarrAPIKey := flag.String("radarr-api-key", "", "API key for Radarr")
	radarrBaseURL := flag.String("radarr-base-url", "", "Base URL for Radarr")
	sonarrAPIKey := flag.String("sonarr-api-key", "", "API key for Sonarr")
	sonarrBaseURL := flag.String("sonarr-base-url", "", "Base URL for Sonarr")

	// Parse the flags
	flag.Parse()

	// Validate that all required flags are provided
	if *overseerAPIKey == "" || *overseerBaseURL == "" ||
		*radarrAPIKey == "" || *radarrBaseURL == "" ||
		*sonarrAPIKey == "" || *sonarrBaseURL == "" {
		log.Fatalf("All API keys and base URLs must be provided via flags")
	}

	// Load the configuration from the provided flags
	config := &Config{
		Overseer: struct {
			APIKey  string
			BaseURL string
		}{
			APIKey:  *overseerAPIKey,
			BaseURL: *overseerBaseURL,
		},
		Radarr: struct {
			APIKey  string
			BaseURL string
		}{
			APIKey:  *radarrAPIKey,
			BaseURL: *radarrBaseURL,
		},
		Sonarr: struct {
			APIKey  string
			BaseURL string
		}{
			APIKey:  *sonarrAPIKey,
			BaseURL: *sonarrBaseURL,
		},
	}

	// Initialize services with the loaded configuration
	overseer := NewOverseerService(config.Overseer.APIKey, config.Overseer.BaseURL)
	radarr := NewRadarrService(config.Radarr.APIKey, config.Radarr.BaseURL)
	sonarr := NewSonarrService(config.Sonarr.APIKey, config.Sonarr.BaseURL)

	overseer.RequestEntireWatchlist()
	radarr.DeleteUnwatched(overseer.WatchListMovies)
	sonarr.DeleteUnwatched(overseer.WatchListSeries)

	// Set up cron job
	c := cron.New()
	c.AddFunc("*/5 * * * *", func() {
		overseer.RequestEntireWatchlist()
		radarr.DeleteUnwatched(overseer.WatchListMovies)
		sonarr.DeleteUnwatched(overseer.WatchListSeries)
	})
	c.Start()

	log.Println("Service started")
	select {}
}
