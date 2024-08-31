package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/robfig/cron/v3"
)

var defaultHttpClient = &http.Client{
	Timeout: 5 * time.Second, // Set a timeout of 5 seconds
}

type ServiceConfig struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

type Config struct {
	Overseer ServiceConfig `yaml:"overseer"`
	Radarr   ServiceConfig `yaml:"radarr"`
	Sonarr   ServiceConfig `yaml:"sonarr"`
}

func main() {
	// Determine the environment
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "development" // Default to development if no environment is set
	}

	// Load the configuration from the appropriate file
	configFile := "config." + env + ".yml"
	config := &Config{}
	loadConfig(configFile, config)

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

func loadConfig(filePath string, config *Config) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("Failed to unmarshal config file: %v", err)
	}
}
