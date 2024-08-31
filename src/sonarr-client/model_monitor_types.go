/*
 * Sonarr
 *
 * Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientSonarr

type MonitorTypes string

// List of MonitorTypes
const (
	UNKNOWN_MonitorTypes            MonitorTypes = "unknown"
	ALL_MonitorTypes                MonitorTypes = "all"
	FUTURE_MonitorTypes             MonitorTypes = "future"
	MISSING_MonitorTypes            MonitorTypes = "missing"
	EXISTING_MonitorTypes           MonitorTypes = "existing"
	FIRST_SEASON_MonitorTypes       MonitorTypes = "firstSeason"
	LAST_SEASON_MonitorTypes        MonitorTypes = "lastSeason"
	LATEST_SEASON_MonitorTypes      MonitorTypes = "latestSeason"
	PILOT_MonitorTypes              MonitorTypes = "pilot"
	RECENT_MonitorTypes             MonitorTypes = "recent"
	MONITOR_SPECIALS_MonitorTypes   MonitorTypes = "monitorSpecials"
	UNMONITOR_SPECIALS_MonitorTypes MonitorTypes = "unmonitorSpecials"
	NONE_MonitorTypes               MonitorTypes = "none"
	SKIP_MonitorTypes               MonitorTypes = "skip"
)