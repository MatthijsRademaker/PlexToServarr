package main

import (
	"context"
	"log"
	"swaggerClientRadarr"
)

type RadarrService struct {
	Servarr   Servarr
	ApiClient *swaggerClientRadarr.APIClient
}

func NewRadarrService(apiKey string, baseUrl string) *RadarrService {
	waitUntilServiceAvailable(baseUrl + "/ping")

	cfg := &swaggerClientRadarr.Configuration{
		BasePath:  baseUrl,
		UserAgent: "Overseer-Client/1.0.0/go",
	}
	context := context.WithValue(context.Background(), swaggerClientRadarr.ContextAPIKey, swaggerClientRadarr.APIKey{Key: apiKey})

	client := swaggerClientRadarr.NewAPIClient(cfg)

	return &RadarrService{
		Servarr: Servarr{
			Context: context,
			Type:    "Radarr",
		},
		ApiClient: client,
	}
}

// Radarr: Handles the deletion and returns the IDs of movies that failed to delete
func (radarr *RadarrService) DeleteUnwatched(moviesToDelete []float64) []float64 {
	var failedDeletions []float64
	for _, movieId := range moviesToDelete {
		log.Printf("Deleting movie with ID: %v", movieId)
		_, err := radarr.ApiClient.MovieApi.ApiV3MovieIdDelete(radarr.Servarr.Context, int32(movieId), nil)
		if err != nil {
			log.Printf("Error deleting movie with ID %v: %v", movieId, err)
			failedDeletions = append(failedDeletions, movieId) // Track failed deletions
		}
	}
	return failedDeletions
}
