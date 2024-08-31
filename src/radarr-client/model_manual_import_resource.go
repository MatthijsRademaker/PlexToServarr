/*
 * Radarr
 *
 * Radarr API docs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientRadarr

type ManualImportResource struct {
	Id                int32                  `json:"id,omitempty"`
	Path              string                 `json:"path,omitempty"`
	RelativePath      string                 `json:"relativePath,omitempty"`
	FolderName        string                 `json:"folderName,omitempty"`
	Name              string                 `json:"name,omitempty"`
	Size              int64                  `json:"size,omitempty"`
	Movie             *MovieResource         `json:"movie,omitempty"`
	Quality           *QualityModel          `json:"quality,omitempty"`
	Languages         []Language             `json:"languages,omitempty"`
	ReleaseGroup      string                 `json:"releaseGroup,omitempty"`
	QualityWeight     int32                  `json:"qualityWeight,omitempty"`
	DownloadId        string                 `json:"downloadId,omitempty"`
	CustomFormats     []CustomFormatResource `json:"customFormats,omitempty"`
	CustomFormatScore int32                  `json:"customFormatScore,omitempty"`
	IndexerFlags      int32                  `json:"indexerFlags,omitempty"`
	Rejections        []Rejection            `json:"rejections,omitempty"`
}