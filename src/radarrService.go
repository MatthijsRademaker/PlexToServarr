package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type RadarrService struct {
	ServarrBase
}

type RadarrMedia struct {
	Title            string          `json:"title"`
	QualityProfileId int             `json:"qualityProfileId"`
	TmdbId           int             `json:"tmdbId"`
	RootFolderPath   string          `json:"rootFolderPath"`
	Monitored        bool            `json:"monitored"`
	AddOptions       map[string]bool `json:"addOptions"`
}

func (base *RadarrService) MapToMediaItem(item PlexItem, id int) RadarrMedia {
	return RadarrMedia{
		Title:            item.Title,
		QualityProfileId: base.QualityProfile.ID,
		TmdbId:           id,
		RootFolderPath:   base.RootFolder,
		Monitored:        true,
		AddOptions: map[string]bool{
			"searchForMovie": true,
		},
	}
}

func (radarr *RadarrService) ProcessWatchList(plexItems []PlexItem) {
	newItems := radarr.GetNewItems(plexItems)
	if len(newItems) == 0 {
		log.Printf("No new movies to process")
		return
	}

	endpoint := fmt.Sprintf("%s/ping?apikey=%s", radarr.Url, radarr.ApiKey)
	waitUntilServiceAvailable(endpoint)
	radarr.SetQualityProfileID()

	media, err := radarr.FetchMedia(plexItems)
	if err != nil {
		log.Fatalf("Failed to fetch media: %v", err)
	}

	err = radarr.AddMedia(media)
	if err != nil {
		log.Fatalf("Failed to add media: %v", err)
	}

	radarr.ProcessedList = append(radarr.ProcessedList, newItems...)
}

func (radarr *RadarrService) QueryDb(item PlexItem, mediaType string) (int, error) {
	encodedApiKey := url.QueryEscape(radarr.DbApiKey)
	encodedTitle := url.QueryEscape(item.Title)
	apiEndpoint := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=%s&query=%s", encodedApiKey, encodedTitle)

	log.Printf("Fetching ID for %s '%s' from %s", mediaType, item.Title, apiEndpoint)
	resp, err := http.Get(apiEndpoint)
	if err != nil {
		log.Printf("Failed to fetch ID for %s '%s': %v", mediaType, item.Title, err)
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch ID for %s '%s'. Status code: %d", mediaType, item.Title, resp.StatusCode)
		return 0, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	var result struct {
		Results []DbResult `json:"results"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("Failed to decode response for %s '%s': %v", mediaType, item.Title, err)
		return 0, err
	}

	log.Printf("Found %d results for %s '%s'", len(result.Results), mediaType, item.Title)
	if len(result.Results) > 0 {
		return result.Results[0].ID, nil
	}

	log.Printf("No ID found for %s '%s'", mediaType, item.Title)
	return 0, fmt.Errorf("no ID found")
}

func (radarr *RadarrService) FetchMedia(items []PlexItem) ([]RadarrMedia, error) {
	mediaItems := make([]RadarrMedia, 0, len(items))
	for _, item := range items {
		id, err := radarr.QueryDb(item, item.Type)
		if err != nil {
			continue
		}

		mediaItems = append(mediaItems, radarr.MapToMediaItem(item, id))
	}
	return mediaItems, nil
}

func (radarr *RadarrService) AddMedia(media []RadarrMedia) error {
	for _, item := range media {
		endpoint := fmt.Sprintf("%s/api/v3/movie?apikey=%s", radarr.Url, radarr.ApiKey)
		err := radarr.uploadMedia(item, endpoint)
		if err != nil {
			return err
		}
	}
	return nil
}
