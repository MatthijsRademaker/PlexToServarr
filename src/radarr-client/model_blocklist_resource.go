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

type BlocklistResource struct {
	Id            int32                  `json:"id,omitempty"`
	MovieId       int32                  `json:"movieId,omitempty"`
	SourceTitle   string                 `json:"sourceTitle,omitempty"`
	Languages     []Language             `json:"languages,omitempty"`
	Quality       *QualityModel          `json:"quality,omitempty"`
	CustomFormats []CustomFormatResource `json:"customFormats,omitempty"`
	Date          time.Time              `json:"date,omitempty"`
	Protocol      *DownloadProtocol      `json:"protocol,omitempty"`
	Indexer       string                 `json:"indexer,omitempty"`
	Message       string                 `json:"message,omitempty"`
	Movie         *MovieResource         `json:"movie,omitempty"`
}