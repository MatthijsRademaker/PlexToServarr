package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type QualityProfile struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ServarrBase struct {
	Url                            string
	ApiKey                         string
	DbApiKey                       string
	RootFolder                     string
	ProcessedList                  []PlexItem
	QualityProfile                 QualityProfile
	WantedQualityProfileFromConfig QualityProfile
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

func (base *ServarrBase) GetNewItems(listToProcess []PlexItem) []PlexItem {
	filterFunc := func(item PlexItem) bool {
		for _, processedItem := range base.ProcessedList {
			if item.Title == processedItem.Title {
				return false
			}
		}
		return true
	}
	return filter(listToProcess, filterFunc)
}

func (servarr *ServarrBase) SetQualityProfileID() {
	url := fmt.Sprintf("%s/api/v3/qualityProfile?apikey=%s", servarr.Url, servarr.ApiKey)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to get quality profiles: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to get quality profiles. Status code: %d", resp.StatusCode)
	}

	var profiles []QualityProfile
	err = json.NewDecoder(resp.Body).Decode(&profiles)
	if err != nil {
		log.Fatalf("Failed to decode quality profiles: %v", err)
	}

	filterFunc := func(profile QualityProfile) bool {
		return profile.Name == servarr.WantedQualityProfileFromConfig.Name
	}

	wantedQualityProfile := filter(profiles, filterFunc)

	if len(wantedQualityProfile) == 0 {
		log.Fatalf("Failed to find quality profile '%s'", servarr.WantedQualityProfileFromConfig.Name)
	}

	servarr.QualityProfile = wantedQualityProfile[0]
}
