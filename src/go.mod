module overwatch

go 1.22.6

require (
	github.com/antihax/optional v1.0.0
	github.com/robfig/cron/v3 v3.0.1
	swaggerClientOverseerr v0.0.0-00010101000000-000000000000
	swaggerClientRadarr v0.0.0-00010101000000-000000000000
	swaggerClientSonarr v0.0.0-00010101000000-000000000000
)

require golang.org/x/oauth2 v0.22.0 // indirect

replace swaggerClientRadarr => ./radarr-client

replace swaggerClientOverseerr => ./overseer-client

replace swaggerClientSonarr => ./sonarr-client
