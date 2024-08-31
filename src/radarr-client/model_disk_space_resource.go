/*
 * Radarr
 *
 * Radarr API docs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientRadarr

type DiskSpaceResource struct {
	Id         int32  `json:"id,omitempty"`
	Path       string `json:"path,omitempty"`
	Label      string `json:"label,omitempty"`
	FreeSpace  int64  `json:"freeSpace,omitempty"`
	TotalSpace int64  `json:"totalSpace,omitempty"`
}