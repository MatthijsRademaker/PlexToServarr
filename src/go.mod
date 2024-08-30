module plex-to-servarr

go 1.22.3

require (
	github.com/robfig/cron/v3 v3.0.1
	swagger v0.0.0-00010101000000-000000000000
)

require (
	github.com/antihax/optional v1.0.0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
)

replace swagger => ./overseer-client
