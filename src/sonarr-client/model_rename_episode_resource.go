/*
 * Sonarr
 *
 * Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientSonarr

type RenameEpisodeResource struct {
	Id             int32   `json:"id,omitempty"`
	SeriesId       int32   `json:"seriesId,omitempty"`
	SeasonNumber   int32   `json:"seasonNumber,omitempty"`
	EpisodeNumbers []int32 `json:"episodeNumbers,omitempty"`
	EpisodeFileId  int32   `json:"episodeFileId,omitempty"`
	ExistingPath   string  `json:"existingPath,omitempty"`
	NewPath        string  `json:"newPath,omitempty"`
}