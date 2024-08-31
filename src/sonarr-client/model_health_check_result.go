/*
 * Sonarr
 *
 * Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientSonarr

type HealthCheckResult string

// List of HealthCheckResult
const (
	OK_HealthCheckResult      HealthCheckResult = "ok"
	NOTICE_HealthCheckResult  HealthCheckResult = "notice"
	WARNING_HealthCheckResult HealthCheckResult = "warning"
	ERROR__HealthCheckResult  HealthCheckResult = "error"
)