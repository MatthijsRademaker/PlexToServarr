package main

import (
	"context"
	"log"
	"swagger"

	"github.com/antihax/optional"
)

type OverseerService struct {
	Context   context.Context
	ApiClient *swagger.APIClient
	Me        swagger.User
}

func NewOverseerService(apiKey string, baseUrl string) *OverseerService {
	waitUntilServiceAvailable(baseUrl + "/status")

	cfg := &swagger.Configuration{
		BasePath:  baseUrl,
		UserAgent: "Overseer-Client/1.0.0/go",
	}
	context := context.WithValue(context.Background(), swagger.ContextAPIKey, swagger.APIKey{Key: apiKey})

	client := swagger.NewAPIClient(cfg)

	me, _, err := client.AuthApi.AuthMeGet(context)

	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	return &OverseerService{
		Context:   context,
		ApiClient: client,
		Me:        me,
	}
}

func (overseer *OverseerService) RequestEntireWatchlist() {

	// Create a map to store already requested media IDs
	requestedMediaIDs := make(map[float64]bool)

	requests, response, err := overseer.ApiClient.RequestApi.RequestGet(overseer.Context, &swagger.RequestApiRequestGetOpts{
		Take: optional.NewFloat64(1000),
	})

	if err != nil {
		log.Printf("Overseerr err fetching requests: %v", err)
		log.Printf("Overseerr response: %v", response)
		return
	}

	// Populate the map with requested media IDs
	for _, request := range requests.Results {
		requestedMediaIDs[request.Media.TmdbId] = true
	}

	// Fetch watchlist
	watchlistPage := float64(1)
	for {
		watchlist, response, err := overseer.ApiClient.SearchApi.DiscoverWatchlistGet(overseer.Context, &swagger.SearchApiDiscoverWatchlistGetOpts{
			Page: optional.NewFloat64(watchlistPage),
		})

		if err != nil {
			log.Printf("Overseerr err fetching watchlist: %v", err)
			log.Printf("Overseerr response: %v", response)
			return
		}

		// Process each media in the watchlist
		for _, media := range watchlist.Results {
			// Check if the media is already requested
			if !requestedMediaIDs[media.TmdbId] {
				overseer.RequestMedia(media)
			} else {
				log.Printf("Media %v already requested, skipping", media.Title)
			}
		}

		// Break the loop if this is the last page
		if watchlistPage >= watchlist.TotalPages {
			break
		}
		watchlistPage++
	}
}

func (overseer *OverseerService) RequestMedia(media swagger.InlineResponse20013Results) {

	switch media.MediaType {
	case "tv":
		overseer.RequestSeries(media)
	case "movie":
		overseer.RequestMovie(media)
	default:
		log.Printf("Unknown media type: %v", media.MediaType)
	}
}

func (overseer *OverseerService) RequestSeries(media swagger.InlineResponse20013Results) {
	log.Printf("Requesting series: %v", media.Title)
	seriesInfo, _, err := overseer.ApiClient.TvApi.TvTvIdGet(overseer.Context, float64(media.TmdbId), &swagger.TvApiTvTvIdGetOpts{})

	if err != nil {
		log.Printf("Overseerr err: %v", err)
		return
	}

	var seasonNumbers []float64 = make([]float64, 0)

	for _, season := range seriesInfo.Seasons {
		seasonNumbers = append(seasonNumbers, season.SeasonNumber)
	}

	_, res, err := overseer.ApiClient.RequestApi.RequestPost(overseer.Context, swagger.RequestBody{
		MediaType:         "tv",
		MediaId:           float64(media.TmdbId),
		TvdbId:            float64(seriesInfo.ExternalIds.TvdbId),
		Is4k:              false,
		ServerId:          float64(0),
		Seasons:           seasonNumbers,
		UserId:            float64(overseer.Me.Id),
		RootFolder:        "/mnt/series",
		LanguageProfileId: float64(1),
		ProfileId:         float64(7),
	})

	if err != nil {
		log.Printf("Overseerr err: %v", err)
		log.Printf("Overseerr response: %v", res)
	}

}

func (overseer *OverseerService) RequestMovie(media swagger.InlineResponse20013Results) {
	log.Printf("Requesting movie: %v", media.Title)

	_, res, err := overseer.ApiClient.RequestApi.RequestPost(overseer.Context, swagger.RequestBody{
		MediaType: "movie",
		MediaId:   float64(media.TmdbId),
		Is4k:      false,
		ServerId:  float64(0),
		UserId:    float64(overseer.Me.Id),
	})

	if err != nil {
		log.Printf("Overseerr err: %v", err)
		log.Printf("Overseerr response: %v", res)
	}
}
