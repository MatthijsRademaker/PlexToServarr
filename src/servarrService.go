package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ServarrBase struct {
	Url        string
	ApiKey     string
	DbApiKey   string
	RootFolder string
}

func (base *ServarrBase) uploadMedia(media interface{}, endpoint string) error {
	body, err := json.Marshal(media)
	if err != nil {
		log.Fatalf("Failed to Marshal payload: %v", err)
	}

	log.Printf("sending request to %s with body %s", endpoint, body)

	resp, err := http.Post(endpoint, "application/json", io.NopCloser(bytes.NewReader(body)))
	if err != nil {
		log.Printf("Failed to add media: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		var errResp []ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err == nil && len(errResp) > 0 {
			log.Printf("Failed to add media. Error: %s", errResp[0].ErrorMessage)
		} else {
			log.Printf("Failed to add media. Status code: %d", resp.StatusCode)
		}
	} else {
		fmt.Printf("Added media successfully.\n")
	}

	return nil
}
