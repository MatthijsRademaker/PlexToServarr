/*
 * Sonarr
 *
 * Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientSonarr

type QualitySource string

// List of QualitySource
const (
	UNKNOWN_QualitySource        QualitySource = "unknown"
	TELEVISION_QualitySource     QualitySource = "television"
	TELEVISION_RAW_QualitySource QualitySource = "televisionRaw"
	WEB_QualitySource            QualitySource = "web"
	WEB_RIP_QualitySource        QualitySource = "webRip"
	DVD_QualitySource            QualitySource = "dvd"
	BLURAY_QualitySource         QualitySource = "bluray"
	BLURAY_RAW_QualitySource     QualitySource = "blurayRaw"
)