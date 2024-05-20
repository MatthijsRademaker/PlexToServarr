package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// TODO make a list of ids so that we can add multiple quality profiles
func (servarr *ServarrBase) getQualityProfileID() int {
	url := fmt.Sprintf("%s/qualityProfile?apikey=%s", servarr.Url, servarr.ApiKey)
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

	for _, profile := range profiles {
		if profile.Name == "HD-1080p" {
			return profile.ID
		}
	}

	log.Fatal("Quality profile 'HD-1080p' not found")
	return 0
}
