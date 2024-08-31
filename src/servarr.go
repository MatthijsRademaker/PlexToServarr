package main

import (
	"context"
	"log"
	"slices"
)

type Servarr struct {
	Context   context.Context
	Watchlist []int32
	Type      string
}

func (servarr *Servarr) GetMediaToDelete(newPlexWatchList []int32) []int32 {

	log.Printf("%v Watchlist: %v", servarr.Type, servarr.Watchlist)
	log.Printf("New Plex Watchlist: %v", newPlexWatchList)

	mediaToDelete := make([]int32, 0)

	for _, servarrMediaId := range servarr.Watchlist {

		if !slices.Contains(newPlexWatchList, servarrMediaId) {
			mediaToDelete = append(mediaToDelete, servarrMediaId)
		}

	}

	log.Printf("%v Media to delete: %v", servarr.Type, mediaToDelete)
	return mediaToDelete
}

// Servarr service with generic delete method
func (servarr *Servarr) DeleteUnwatched(newPlexWatchList []int32, deleteFunc func(int32) error) {
	seriesToDelete := servarr.GetMediaToDelete(newPlexWatchList)

	for _, id := range seriesToDelete {
		// Use the provided delete function to delete the item
		err := deleteFunc(id)
		if err != nil {
			// Handle the error if needed
		} else {
			servarr.removeFromWatchlistCache(id)
		}
	}

	servarr.UpdateWatchList(newPlexWatchList)
}

func (servarr *Servarr) removeFromWatchlistCache(id int32) {
	for i, value := range servarr.Watchlist {
		if value == id {
			servarr.Watchlist = append(servarr.Watchlist[:i], servarr.Watchlist[i+1:]...)
			break
		}
	}
}

func (servarr *Servarr) UpdateWatchList(newPlexWatchList []int32) {

	// Just add new items from newPlexWatchList if not already present
	for _, value := range newPlexWatchList {
		if !slices.Contains(servarr.Watchlist, value) {
			servarr.Watchlist = append(servarr.Watchlist, value)
		}
	}
	log.Printf("Updated %v watchlist: %v", servarr.Type, servarr.Watchlist)
}
