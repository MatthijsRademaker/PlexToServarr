package main

import (
	"context"
	"log"
	"swaggerClientSonarr"
)

type SonarrService struct {
	Servarr   Servarr
	ApiClient *swaggerClientSonarr.APIClient
}

func NewSonarrService(apiKey string, baseUrl string) *SonarrService {
	waitUntilServiceAvailable(baseUrl + "/ping")

	cfg := &swaggerClientSonarr.Configuration{
		BasePath:  baseUrl,
		UserAgent: "Overseer-Client/1.0.0/go",
	}
	context := context.WithValue(context.Background(), swaggerClientSonarr.ContextAPIKey, swaggerClientSonarr.APIKey{Key: apiKey})

	client := swaggerClientSonarr.NewAPIClient(cfg)

	return &SonarrService{
		Servarr: Servarr{
			Context: context,
			Type:    "Sonarr",
		},
		ApiClient: client,
	}
}

// Sonarr: Handles the deletion and returns the IDs of series that failed to delete
func (sonarr *SonarrService) DeleteUnwatched(seriesToDelete []float64) []float64 {
	var failedDeletions []float64
	for _, seriesId := range seriesToDelete {
		log.Printf("Deleting series with ID: %v", seriesId)
		_, err := sonarr.ApiClient.SeriesApi.ApiV3SeriesIdDelete(sonarr.Servarr.Context, int32(seriesId), nil)
		if err != nil {
			log.Printf("Error deleting series with ID %v: %v", seriesId, err)
			failedDeletions = append(failedDeletions, seriesId) // Track failed deletions
		}
	}
	return failedDeletions
}

func (overseer *OverseerService) ProcessDeletions(radarr *RadarrService, sonarr *SonarrService) {
	// Delete unwatched movies and update the list of MoviesToDelete
	failedMovies := radarr.DeleteUnwatched(overseer.MoviesToDelete)
	overseer.MoviesToDelete = failedMovies

	// Delete unwatched series and update the list of SeriesToDelete
	failedSeries := sonarr.DeleteUnwatched(overseer.SeriesToDelete)
	overseer.SeriesToDelete = failedSeries
}
