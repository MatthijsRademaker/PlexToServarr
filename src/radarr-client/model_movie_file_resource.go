/*
 * Radarr
 *
 * Radarr API docs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientRadarr

import (
	"time"
)

type MovieFileResource struct {
	Id                  int32                  `json:"id,omitempty"`
	MovieId             int32                  `json:"movieId,omitempty"`
	RelativePath        string                 `json:"relativePath,omitempty"`
	Path                string                 `json:"path,omitempty"`
	Size                int64                  `json:"size,omitempty"`
	DateAdded           time.Time              `json:"dateAdded,omitempty"`
	SceneName           string                 `json:"sceneName,omitempty"`
	ReleaseGroup        string                 `json:"releaseGroup,omitempty"`
	Edition             string                 `json:"edition,omitempty"`
	Languages           []Language             `json:"languages,omitempty"`
	Quality             *QualityModel          `json:"quality,omitempty"`
	CustomFormats       []CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore   int32                  `json:"customFormatScore,omitempty"`
	IndexerFlags        int32                  `json:"indexerFlags,omitempty"`
	MediaInfo           *MediaInfoResource     `json:"mediaInfo,omitempty"`
	OriginalFilePath    string                 `json:"originalFilePath,omitempty"`
	QualityCutoffNotMet bool                   `json:"qualityCutoffNotMet,omitempty"`
}