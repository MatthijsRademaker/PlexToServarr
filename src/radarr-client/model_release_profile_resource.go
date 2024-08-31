/*
 * Radarr
 *
 * Radarr API docs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientRadarr

type ReleaseProfileResource struct {
	Id        int32   `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	Enabled   bool    `json:"enabled,omitempty"`
	IndexerId int32   `json:"indexerId,omitempty"`
	Tags      []int32 `json:"tags,omitempty"`
}