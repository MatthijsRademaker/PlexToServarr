package main

import (
	"context"
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
			Context:   context,
			Watchlist: make([]int32, 0),
			Type:      "Sonarr",
		},
		ApiClient: client,
	}
}

func (sonarr *SonarrService) DeleteUnwatched(newPlexWatchList []int32) {
	sonarr.Servarr.DeleteUnwatched(newPlexWatchList, func(id int32) error {
		_, err := sonarr.ApiClient.SeriesApi.ApiV3SeriesIdDelete(sonarr.Servarr.Context, id, nil)
		return err
	})
}
