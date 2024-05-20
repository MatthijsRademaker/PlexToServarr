package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// filter function to return only even numbers from a slice
func filter[T any](slice []T, condition func(T) bool) []T {
	var result []T
	for _, value := range slice {
		if condition(value) {
			result = append(result, value)
		}
	}
	return result
}

func httpPost[T any](body []byte, url string, result *T) error {

	resp, err := http.Post(url, "application/json", io.NopCloser(bytes.NewReader(body)))
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("Failed to decode response for %s: %v", url, err)
		return err
	}
	return nil
}
