package main

import (
	"log"
	"os"

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
	overseerAPIKey := os.Getenv("OVERSEER_API_KEY")
	overseerBaseURL := os.Getenv("OVERSEER_BASE_URL")
	radarrAPIKey := os.Getenv("RADARR_API_KEY")
	radarrBaseURL := os.Getenv("RADARR_BASE_URL")
	sonarrAPIKey := os.Getenv("SONARR_API_KEY")
	sonarrBaseURL := os.Getenv("SONARR_BASE_URL")

	// Validate that all required environment variables are provided
	if overseerAPIKey == "" || overseerBaseURL == "" ||
		radarrAPIKey == "" || radarrBaseURL == "" ||
		sonarrAPIKey == "" || sonarrBaseURL == "" {
		log.Fatalf("All API keys and base URLs must be provided via environment variables")
	}

	// Load the configuration from the environment variables
	config := &Config{
		Overseer: struct {
			APIKey  string
			BaseURL string
		}{
			APIKey:  overseerAPIKey,
			BaseURL: overseerBaseURL,
		},
		Radarr: struct {
			APIKey  string
			BaseURL string
		}{
			APIKey:  radarrAPIKey,
			BaseURL: radarrBaseURL,
		},
		Sonarr: struct {
			APIKey  string
			BaseURL string
		}{
			APIKey:  sonarrAPIKey,
			BaseURL: sonarrBaseURL,
		},
	}

	// Initialize services with the loaded configuration
	overseer := NewOverseerService(config.Overseer.APIKey, config.Overseer.BaseURL)
	radarr := NewRadarrService(config.Radarr.APIKey, config.Radarr.BaseURL)
	sonarr := NewSonarrService(config.Sonarr.APIKey, config.Sonarr.BaseURL)

	overseer.RequestEntireWatchlist(true)

	// Set up cron job
	c := cron.New()
	c.AddFunc("*/5 * * * *", func() {
		overseer.RequestEntireWatchlist(false)
		overseer.ProcessDeletions(radarr, sonarr)
	})
	c.Start()

	log.Println("Service started")
	select {}
}
