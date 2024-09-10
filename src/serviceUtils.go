package main

import (
	"log"
	"net/http"
	"os"
	"slices"
	"time"
)

var defaultHttpClient = &http.Client{
	Timeout: 5 * time.Second, // Set a timeout of 5 seconds
}

func waitUntilServiceAvailable(url string) bool {
	log.Println("Waiting for service to become available:", url)
	for {
		if isServiceAvailable(url) {
			return true
		}

		time.Sleep(10 * time.Second)
	}
}

func isServiceAvailable(url string) bool {
	log.Println("Checking availability of:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return false
	}

	resp, err := defaultHttpClient.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			log.Printf("Timeout error: %v", err)
		} else {
			log.Printf("Request error: %v", err)
		}
		return false
	}
	defer resp.Body.Close()

	log.Println("Service available. Status Code:", resp.StatusCode)
	return true
}

// Helper function to check if a value exists in a slice
func AddToWatchList(slice []int32, value int32) []int32 {
	if !slices.Contains(slice, value) {
		return append(slice, value)
	}
	return slice
}
