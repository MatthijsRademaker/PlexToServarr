package main

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Radarr RadarrService `json:"radarr"`
	Sonarr SonarrService `json:"sonarr"`
	Plex   struct {
		URL    string `json:"url"`
		ApiKey string `json:"apiKey"`
	} `json:"plex"`
}

// Function to load configuration from a file
func loadConfig(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()

	bytes, err := io.ReadAll(configFile)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	return config, err
}
