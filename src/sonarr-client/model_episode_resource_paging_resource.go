/*
 * Sonarr
 *
 * Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientSonarr

type EpisodeResourcePagingResource struct {
	Page          int32             `json:"page,omitempty"`
	PageSize      int32             `json:"pageSize,omitempty"`
	SortKey       string            `json:"sortKey,omitempty"`
	SortDirection *SortDirection    `json:"sortDirection,omitempty"`
	TotalRecords  int32             `json:"totalRecords,omitempty"`
	Records       []EpisodeResource `json:"records,omitempty"`
}
