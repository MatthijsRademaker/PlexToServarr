package main

import (
	"context"
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
			Context:   context,
			Watchlist: make([]int32, 0),
			Type:      "Radarr",
		},
		ApiClient: client,
	}
}

func (radarr *RadarrService) DeleteUnwatched(newPlexWatchList []int32) {
	radarr.Servarr.DeleteUnwatched(newPlexWatchList, func(id int32) error {
		_, err := radarr.ApiClient.MovieApi.ApiV3MovieIdDelete(radarr.Servarr.Context, id, nil)
		return err
	})
}
