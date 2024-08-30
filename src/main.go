package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/robfig/cron/v3"
)

var defaultHttpClient = &http.Client{
	Timeout: 5 * time.Second, // Set a timeout of 5 seconds
}

func main() {
	// Define command-line flags
	apiKey := flag.String("api-key", "", "API key for Overseer")
	baseURL := flag.String("base-url", "", "Base URL for Overseer API")
	cronSchedule := flag.String("cron", "*/5 * * * *", "Cron schedule for running the job")

	// Parse the flags
	flag.Parse()

	// Check if required flags are provided
	if *apiKey == "" || *baseURL == "" {
		log.Fatal("API key and base URL are required. Use -api-key and -base-url flags.")
	}

	// Initialize OverseerService with command-line arguments
	overseer := NewOverseerService(*apiKey, *baseURL)

	// Initial run
	overseer.RequestEntireWatchlist()

	// Set up cron job
	c := cron.New()
	c.AddFunc(*cronSchedule, func() {
		overseer.RequestEntireWatchlist()
	})
	c.Start()

	log.Println("Service started")
	select {}
}
