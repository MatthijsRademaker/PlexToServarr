package main

import (
	"context"
	"log"
	"swaggerClientOverseerr"

	"github.com/antihax/optional"
)

type OverseerService struct {
	Context                 context.Context
	ApiClient               *swaggerClientOverseerr.APIClient
	Me                      swaggerClientOverseerr.User
	PreviousWatchListMovies []float64
	PreviousWatchListSeries []float64
	MoviesToDelete          []float64
	SeriesToDelete          []float64
}

func NewOverseerService(apiKey string, baseUrl string) *OverseerService {
	waitUntilServiceAvailable(baseUrl + "/status")

	cfg := &swaggerClientOverseerr.Configuration{
		BasePath:  baseUrl + "/api/v1",
		UserAgent: "Overseer-Client/1.0.0/go",
	}
	context := context.WithValue(context.Background(), swaggerClientOverseerr.ContextAPIKey, swaggerClientOverseerr.APIKey{Key: apiKey})

	client := swaggerClientOverseerr.NewAPIClient(cfg)

	me, _, err := client.AuthApi.AuthMeGet(context)

	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	return &OverseerService{
		Context:                 context,
		ApiClient:               client,
		Me:                      me,
		PreviousWatchListMovies: make([]float64, 0),
		PreviousWatchListSeries: make([]float64, 0),
		MoviesToDelete:          make([]float64, 0),
		SeriesToDelete:          make([]float64, 0),
	}
}

func (overseer *OverseerService) RequestEntireWatchlist() {
	watchlistPage := float64(1)
	var currentWatchListMovies []float64
	var currentWatchListSeries []float64

	for {
		watchlist, response, err := overseer.ApiClient.SearchApi.DiscoverWatchlistGet(overseer.Context, &swaggerClientOverseerr.SearchApiDiscoverWatchlistGetOpts{
			Page: optional.NewFloat64(watchlistPage),
		})

		if err != nil {
			log.Printf("Overseerr err fetching watchlist: %v", err)
			log.Printf("Overseerr response: %v", response)
			return
		}

		// Process each media in the watchlist
		for _, media := range watchlist.Results {
			overseer.ProcessWatchlistItem(media, &currentWatchListMovies, &currentWatchListSeries)
		}

		// Break the loop if this is the last page
		if watchlistPage >= watchlist.TotalPages {
			break
		}
		watchlistPage++
	}

	// Compare with the previous watchlist state and handle removed items
	overseer.HandleRemovedItems(currentWatchListMovies, currentWatchListSeries)

	// Update the previous watchlist state
	overseer.PreviousWatchListMovies = currentWatchListMovies
	overseer.PreviousWatchListSeries = currentWatchListSeries
}

// ProcessWatchlistItem checks and processes a single media item in the watchlist
func (overseer *OverseerService) ProcessWatchlistItem(media swaggerClientOverseerr.InlineResponse20013Results, currentMovies *[]float64, currentSeries *[]float64) {
	switch media.MediaType {
	case "tv":
		overseer.RequestSeries(media, currentSeries)
	case "movie":
		*currentMovies = append(*currentMovies, media.TmdbId)
		if !ItemInList(overseer.PreviousWatchListMovies, media.TmdbId) {
			overseer.RequestMovie(media)
		} else {
			log.Printf("Movie %v already requested, skipping", media.Title)
		}
	}
}

// HandleRemovedItems compares the current and previous watchlists and removes items that are no longer in the current list
func (overseer *OverseerService) HandleRemovedItems(currentMovies []float64, currentSeries []float64) {
	removedMovies := FindRemovedItems(overseer.PreviousWatchListMovies, currentMovies)
	removedSeries := FindRemovedItems(overseer.PreviousWatchListSeries, currentSeries)

	for _, movie := range removedMovies {
		log.Printf("Movie %v is no longer in the watchlist, removing from requested list", movie)
		overseer.MoviesToDelete = append(overseer.MoviesToDelete, movie)
		// Add logic to unsubscribe from services if needed
	}

	for _, series := range removedSeries {
		log.Printf("Series %v is no longer in the watchlist, removing from requested list", series)
		overseer.SeriesToDelete = append(overseer.SeriesToDelete, series)
		// Add logic to unsubscribe from services if needed
	}
}

// FindRemovedItems identifies items that are in the previous list but not in the current one
func FindRemovedItems(previousList, currentList []float64) []float64 {
	removed := make([]float64, 0)
	previousMap := make(map[float64]bool)

	for _, item := range previousList {
		previousMap[item] = true
	}

	for _, item := range currentList {
		delete(previousMap, item)
	}

	for item := range previousMap {
		removed = append(removed, item)
	}

	return removed
}

// ItemInList checks if an item is in a list
func ItemInList(list []float64, item float64) bool {
	for _, listItem := range list {
		if listItem == item {
			return true
		}
	}
	return false
}

func (overseer *OverseerService) RequestSeries(media swaggerClientOverseerr.InlineResponse20013Results, currentSeries *[]float64) {
	log.Printf("Requesting series: %v", media.Title)
	seriesInfo, _, err := overseer.ApiClient.TvApi.TvTvIdGet(overseer.Context, media.TmdbId, &swaggerClientOverseerr.TvApiTvTvIdGetOpts{})

	if err != nil {
		log.Printf("Overseerr err: %v", err)
		return
	}

	*currentSeries = append(*currentSeries, seriesInfo.ExternalIds.TvdbId)

	if ItemInList(overseer.PreviousWatchListSeries, seriesInfo.ExternalIds.TvdbId) {
		log.Printf("TV series %v already requested, skipping", media.Title)
		return
	}

	var seasonNumbers []float64 = make([]float64, 0)

	for _, season := range seriesInfo.Seasons {
		if season.SeasonNumber == 0 {
			continue
		}
		seasonNumbers = append(seasonNumbers, season.SeasonNumber)
	}

	_, res, err := overseer.ApiClient.RequestApi.RequestPost(overseer.Context, swaggerClientOverseerr.RequestBody{
		MediaType:         "tv",
		MediaId:           media.TmdbId,
		TvdbId:            seriesInfo.ExternalIds.TvdbId,
		Is4k:              false,
		ServerId:          float64(0),
		Seasons:           seasonNumbers,
		UserId:            float64(overseer.Me.Id),
		RootFolder:        "/data/media/tv",
		LanguageProfileId: float64(1),
		ProfileId:         float64(7),
	})

	if err != nil {
		log.Printf("Overseerr err: %v", err)
		log.Printf("Overseerr response: %v", res)
	}
}

func (overseer *OverseerService) RequestMovie(media swaggerClientOverseerr.InlineResponse20013Results) {
	log.Printf("Requesting movie: %v", media.Title)

	_, res, err := overseer.ApiClient.RequestApi.RequestPost(overseer.Context, swaggerClientOverseerr.RequestBody{
		MediaType: "movie",
		MediaId:   media.TmdbId,
		Is4k:      false,
		ServerId:  float64(0),
		UserId:    float64(overseer.Me.Id),
	})

	if err != nil {
		log.Printf("Overseerr err: %v", err)
		log.Printf("Overseerr response: %v", res)
	}
}
