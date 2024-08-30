package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

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
