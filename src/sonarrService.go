package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type SonarrMedia struct {
	Title            string          `json:"title"`
	QualityProfileId int             `json:"qualityProfileId"`
	TvdbId           int             `json:"tvdbId"`
	RootFolderPath   string          `json:"rootFolderPath"`
	Monitored        bool            `json:"monitored"`
	AddOptions       map[string]bool `json:"addOptions"`
}

type SonarrService struct {
	ServarrBase
}

func (base *SonarrService) MapToMediaItem(item PlexItem, id int) SonarrMedia {
	return SonarrMedia{
		Title:            item.Title,
		QualityProfileId: base.QualityProfile.ID,
		TvdbId:           id,
		RootFolderPath:   base.RootFolder,
		Monitored:        true,
		AddOptions: map[string]bool{
			"searchForMovie": true,
		},
	}
}

func (sonarr *SonarrService) ProcessWatchList(plexItems []PlexItem) {

	newItems := sonarr.GetNewItems(plexItems)
	if len(newItems) == 0 {
		log.Printf("No new shows to process")
		return
	}

	endpoint := fmt.Sprintf("%s/ping?apikey=%s", sonarr.Url, sonarr.ApiKey)
	waitUntilServiceAvailable(endpoint)
	sonarr.SetQualityProfileID()

	media, err := sonarr.FetchMedia(plexItems)
	if err != nil {
		log.Fatalf("Failed to fetch media: %v", err)
	}

	err = sonarr.AddMedia(media)
	if err != nil {
		log.Fatalf("Failed to add media: %v", err)
	}

	sonarr.ProcessedList = append(sonarr.ProcessedList, newItems...)
}

func (sonarr *SonarrService) QueryDb(item PlexItem, mediaType string) (int, error) {
	theTvDbUrl := "https://api4.thetvdb.com/v4"
	body, err := json.Marshal(map[string]string{
		"apikey": sonarr.DbApiKey,
	})

	apiEndpoint := fmt.Sprintf("%s/login", theTvDbUrl)

	var loginResult struct {
		Status string `json:"status"`
		Data   struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	err = httpPost(body, apiEndpoint, &loginResult)
	if err != nil {
		log.Printf("Failed to login to TheTVDB: %v", err)
		return 0, err
	}

	encodedTitle := url.QueryEscape(item.Title)
	apiEndpoint = fmt.Sprintf("%s/search?query=%s&type=series", theTvDbUrl, encodedTitle)
	log.Printf("Fetching ID for %s '%s' from %s", mediaType, item.Title, apiEndpoint)
	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Add the Authorization header
	req.Header.Add("Authorization", "Bearer "+loginResult.Data.Token)

	// Create an HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to fetch ID for %s '%s'. Status code: %d", mediaType, item.Title, resp.StatusCode)
		return 0, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	var result struct {
		Data []struct {
			TvdbId string `json:"tvdb_id"`
		} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("Failed to decode response for %s '%s': %v", mediaType, item.Title, err)
		return 0, err
	}

	if len(result.Data) > 0 {
		id, err := strconv.Atoi(result.Data[0].TvdbId)
		if err != nil {
			// ... handle error
			panic(err)
		}
		return id, nil
	}

	log.Printf("No ID found for %s '%s'", mediaType, item.Title)
	return 0, fmt.Errorf("no ID found")
}

func (sonarr *SonarrService) FetchMedia(items []PlexItem) ([]SonarrMedia, error) {
	mediaItems := make([]SonarrMedia, 0, len(items))
	for _, item := range items {
		id, err := sonarr.QueryDb(item, item.Type)
		if err != nil {
			continue
		}

		mediaItems = append(mediaItems, sonarr.MapToMediaItem(item, id))
	}
	return mediaItems, nil
}

func (sonarr *SonarrService) AddMedia(media []SonarrMedia) error {
	for _, item := range media {
		endpoint := fmt.Sprintf("%s/api/v3/series?apikey=%s", sonarr.Url, sonarr.ApiKey)
		err := sonarr.uploadMedia(item, endpoint)
		if err != nil {
			return err
		}
	}
	return nil
}
