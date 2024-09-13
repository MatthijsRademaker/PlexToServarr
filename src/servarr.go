package main

import (
	"context"
)

type Servarr struct {
	Context context.Context
	Type    string
}

// Servarr service with generic delete method
func (servarr *Servarr) DeleteUnwatched(idsToDelete []float64, deleteFunc func(float64) error) {

	for _, id := range idsToDelete {
		// Use the provided delete function to delete the item
		err := deleteFunc(id)
		if err != nil {
			// Handle the error if needed
		}
	}
}
