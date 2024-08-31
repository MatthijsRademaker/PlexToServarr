/*
 * Sonarr
 *
 * Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientSonarr

type ImportListExclusionResource struct {
	Id     int32  `json:"id,omitempty"`
	TvdbId int32  `json:"tvdbId,omitempty"`
	Title  string `json:"title,omitempty"`
}