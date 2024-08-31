/*
 * Radarr
 *
 * Radarr API docs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientRadarr

type MonitorTypes string

// List of MonitorTypes
const (
	MOVIE_ONLY_MonitorTypes           MonitorTypes = "movieOnly"
	MOVIE_AND_COLLECTION_MonitorTypes MonitorTypes = "movieAndCollection"
	NONE_MonitorTypes                 MonitorTypes = "none"
)