# Go API client for swagger

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *{server}/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthApi* | [**AuthLocalPost**](docs/AuthApi.md#authlocalpost) | **Post** /auth/local | Sign in using a local account
*AuthApi* | [**AuthLogoutPost**](docs/AuthApi.md#authlogoutpost) | **Post** /auth/logout | Sign out and clear session cookie
*AuthApi* | [**AuthMeGet**](docs/AuthApi.md#authmeget) | **Get** /auth/me | Get logged-in user
*AuthApi* | [**AuthPlexPost**](docs/AuthApi.md#authplexpost) | **Post** /auth/plex | Sign in using a Plex token
*CollectionApi* | [**CollectionCollectionIdGet**](docs/CollectionApi.md#collectioncollectionidget) | **Get** /collection/{collectionId} | Get collection details
*IssueApi* | [**IssueCommentCommentIdDelete**](docs/IssueApi.md#issuecommentcommentiddelete) | **Delete** /issueComment/{commentId} | Delete issue comment
*IssueApi* | [**IssueCommentCommentIdGet**](docs/IssueApi.md#issuecommentcommentidget) | **Get** /issueComment/{commentId} | Get issue comment
*IssueApi* | [**IssueCommentCommentIdPut**](docs/IssueApi.md#issuecommentcommentidput) | **Put** /issueComment/{commentId} | Update issue comment
*IssueApi* | [**IssueCountGet**](docs/IssueApi.md#issuecountget) | **Get** /issue/count | Gets issue counts
*IssueApi* | [**IssueGet**](docs/IssueApi.md#issueget) | **Get** /issue | Get all issues
*IssueApi* | [**IssueIssueIdCommentPost**](docs/IssueApi.md#issueissueidcommentpost) | **Post** /issue/{issueId}/comment | Create a comment
*IssueApi* | [**IssueIssueIdDelete**](docs/IssueApi.md#issueissueiddelete) | **Delete** /issue/{issueId} | Delete issue
*IssueApi* | [**IssueIssueIdGet**](docs/IssueApi.md#issueissueidget) | **Get** /issue/{issueId} | Get issue
*IssueApi* | [**IssueIssueIdStatusPost**](docs/IssueApi.md#issueissueidstatuspost) | **Post** /issue/{issueId}/{status} | Update an issue&#x27;s status
*IssueApi* | [**IssuePost**](docs/IssueApi.md#issuepost) | **Post** /issue | Create new issue
*MediaApi* | [**MediaGet**](docs/MediaApi.md#mediaget) | **Get** /media | Get media
*MediaApi* | [**MediaMediaIdDelete**](docs/MediaApi.md#mediamediaiddelete) | **Delete** /media/{mediaId} | Delete media item
*MediaApi* | [**MediaMediaIdStatusPost**](docs/MediaApi.md#mediamediaidstatuspost) | **Post** /media/{mediaId}/{status} | Update media status
*MediaApi* | [**MediaMediaIdWatchDataGet**](docs/MediaApi.md#mediamediaidwatchdataget) | **Get** /media/{mediaId}/watch_data | Get watch data
*MoviesApi* | [**MovieMovieIdGet**](docs/MoviesApi.md#moviemovieidget) | **Get** /movie/{movieId} | Get movie details
*MoviesApi* | [**MovieMovieIdRatingsGet**](docs/MoviesApi.md#moviemovieidratingsget) | **Get** /movie/{movieId}/ratings | Get movie ratings
*MoviesApi* | [**MovieMovieIdRatingscombinedGet**](docs/MoviesApi.md#moviemovieidratingscombinedget) | **Get** /movie/{movieId}/ratingscombined | Get RT and IMDB movie ratings combined
*MoviesApi* | [**MovieMovieIdRecommendationsGet**](docs/MoviesApi.md#moviemovieidrecommendationsget) | **Get** /movie/{movieId}/recommendations | Get recommended movies
*MoviesApi* | [**MovieMovieIdSimilarGet**](docs/MoviesApi.md#moviemovieidsimilarget) | **Get** /movie/{movieId}/similar | Get similar movies
*OtherApi* | [**KeywordKeywordIdGet**](docs/OtherApi.md#keywordkeywordidget) | **Get** /keyword/{keywordId} | Get keyword
*OtherApi* | [**WatchprovidersMoviesGet**](docs/OtherApi.md#watchprovidersmoviesget) | **Get** /watchproviders/movies | Get watch provider movies
*OtherApi* | [**WatchprovidersRegionsGet**](docs/OtherApi.md#watchprovidersregionsget) | **Get** /watchproviders/regions | Get watch provider regions
*OtherApi* | [**WatchprovidersTvGet**](docs/OtherApi.md#watchproviderstvget) | **Get** /watchproviders/tv | Get watch provider series
*PersonApi* | [**PersonPersonIdCombinedCreditsGet**](docs/PersonApi.md#personpersonidcombinedcreditsget) | **Get** /person/{personId}/combined_credits | Get combined credits
*PersonApi* | [**PersonPersonIdGet**](docs/PersonApi.md#personpersonidget) | **Get** /person/{personId} | Get person details
*PublicApi* | [**StatusAppdataGet**](docs/PublicApi.md#statusappdataget) | **Get** /status/appdata | Get application data volume status
*PublicApi* | [**StatusGet**](docs/PublicApi.md#statusget) | **Get** /status | Get Overseerr status
*RequestApi* | [**RequestCountGet**](docs/RequestApi.md#requestcountget) | **Get** /request/count | Gets request counts
*RequestApi* | [**RequestGet**](docs/RequestApi.md#requestget) | **Get** /request | Get all requests
*RequestApi* | [**RequestPost**](docs/RequestApi.md#requestpost) | **Post** /request | Create new request
*RequestApi* | [**RequestRequestIdDelete**](docs/RequestApi.md#requestrequestiddelete) | **Delete** /request/{requestId} | Delete request
*RequestApi* | [**RequestRequestIdGet**](docs/RequestApi.md#requestrequestidget) | **Get** /request/{requestId} | Get MediaRequest
*RequestApi* | [**RequestRequestIdPut**](docs/RequestApi.md#requestrequestidput) | **Put** /request/{requestId} | Update MediaRequest
*RequestApi* | [**RequestRequestIdRetryPost**](docs/RequestApi.md#requestrequestidretrypost) | **Post** /request/{requestId}/retry | Retry failed request
*RequestApi* | [**RequestRequestIdStatusPost**](docs/RequestApi.md#requestrequestidstatuspost) | **Post** /request/{requestId}/{status} | Update a request&#x27;s status
*SearchApi* | [**DiscoverGenresliderMovieGet**](docs/SearchApi.md#discovergenreslidermovieget) | **Get** /discover/genreslider/movie | Get genre slider data for movies
*SearchApi* | [**DiscoverGenresliderTvGet**](docs/SearchApi.md#discovergenreslidertvget) | **Get** /discover/genreslider/tv | Get genre slider data for TV series
*SearchApi* | [**DiscoverKeywordKeywordIdMoviesGet**](docs/SearchApi.md#discoverkeywordkeywordidmoviesget) | **Get** /discover/keyword/{keywordId}/movies | Get movies from keyword
*SearchApi* | [**DiscoverMoviesGenreGenreIdGet**](docs/SearchApi.md#discovermoviesgenregenreidget) | **Get** /discover/movies/genre/{genreId} | Discover movies by genre
*SearchApi* | [**DiscoverMoviesGet**](docs/SearchApi.md#discovermoviesget) | **Get** /discover/movies | Discover movies
*SearchApi* | [**DiscoverMoviesLanguageLanguageGet**](docs/SearchApi.md#discovermovieslanguagelanguageget) | **Get** /discover/movies/language/{language} | Discover movies by original language
*SearchApi* | [**DiscoverMoviesStudioStudioIdGet**](docs/SearchApi.md#discovermoviesstudiostudioidget) | **Get** /discover/movies/studio/{studioId} | Discover movies by studio
*SearchApi* | [**DiscoverMoviesUpcomingGet**](docs/SearchApi.md#discovermoviesupcomingget) | **Get** /discover/movies/upcoming | Upcoming movies
*SearchApi* | [**DiscoverTrendingGet**](docs/SearchApi.md#discovertrendingget) | **Get** /discover/trending | Trending movies and TV
*SearchApi* | [**DiscoverTvGenreGenreIdGet**](docs/SearchApi.md#discovertvgenregenreidget) | **Get** /discover/tv/genre/{genreId} | Discover TV shows by genre
*SearchApi* | [**DiscoverTvGet**](docs/SearchApi.md#discovertvget) | **Get** /discover/tv | Discover TV shows
*SearchApi* | [**DiscoverTvLanguageLanguageGet**](docs/SearchApi.md#discovertvlanguagelanguageget) | **Get** /discover/tv/language/{language} | Discover TV shows by original language
*SearchApi* | [**DiscoverTvNetworkNetworkIdGet**](docs/SearchApi.md#discovertvnetworknetworkidget) | **Get** /discover/tv/network/{networkId} | Discover TV shows by network
*SearchApi* | [**DiscoverTvUpcomingGet**](docs/SearchApi.md#discovertvupcomingget) | **Get** /discover/tv/upcoming | Discover Upcoming TV shows
*SearchApi* | [**DiscoverWatchlistGet**](docs/SearchApi.md#discoverwatchlistget) | **Get** /discover/watchlist | Get the Plex watchlist.
*SearchApi* | [**SearchCompanyGet**](docs/SearchApi.md#searchcompanyget) | **Get** /search/company | Search for companies
*SearchApi* | [**SearchGet**](docs/SearchApi.md#searchget) | **Get** /search | Search for movies, TV shows, or people
*SearchApi* | [**SearchKeywordGet**](docs/SearchApi.md#searchkeywordget) | **Get** /search/keyword | Search for keywords
*ServiceApi* | [**ServiceRadarrGet**](docs/ServiceApi.md#serviceradarrget) | **Get** /service/radarr | Get non-sensitive Radarr server list
*ServiceApi* | [**ServiceRadarrRadarrIdGet**](docs/ServiceApi.md#serviceradarrradarridget) | **Get** /service/radarr/{radarrId} | Get Radarr server quality profiles and root folders
*ServiceApi* | [**ServiceSonarrGet**](docs/ServiceApi.md#servicesonarrget) | **Get** /service/sonarr | Get non-sensitive Sonarr server list
*ServiceApi* | [**ServiceSonarrLookupTmdbIdGet**](docs/ServiceApi.md#servicesonarrlookuptmdbidget) | **Get** /service/sonarr/lookup/{tmdbId} | Get series from Sonarr
*ServiceApi* | [**ServiceSonarrSonarrIdGet**](docs/ServiceApi.md#servicesonarrsonarridget) | **Get** /service/sonarr/{sonarrId} | Get Sonarr server quality profiles and root folders
*SettingsApi* | [**SettingsAboutGet**](docs/SettingsApi.md#settingsaboutget) | **Get** /settings/about | Get server stats
*SettingsApi* | [**SettingsCacheCacheIdFlushPost**](docs/SettingsApi.md#settingscachecacheidflushpost) | **Post** /settings/cache/{cacheId}/flush | Flush a specific cache
*SettingsApi* | [**SettingsCacheGet**](docs/SettingsApi.md#settingscacheget) | **Get** /settings/cache | Get a list of active caches
*SettingsApi* | [**SettingsDiscoverAddPost**](docs/SettingsApi.md#settingsdiscoveraddpost) | **Post** /settings/discover/add | Add a new slider
*SettingsApi* | [**SettingsDiscoverGet**](docs/SettingsApi.md#settingsdiscoverget) | **Get** /settings/discover | Get all discover sliders
*SettingsApi* | [**SettingsDiscoverPost**](docs/SettingsApi.md#settingsdiscoverpost) | **Post** /settings/discover | Batch update all sliders.
*SettingsApi* | [**SettingsDiscoverResetGet**](docs/SettingsApi.md#settingsdiscoverresetget) | **Get** /settings/discover/reset | Reset all discover sliders
*SettingsApi* | [**SettingsDiscoverSliderIdDelete**](docs/SettingsApi.md#settingsdiscoverslideriddelete) | **Delete** /settings/discover/{sliderId} | Delete slider by ID
*SettingsApi* | [**SettingsDiscoverSliderIdPut**](docs/SettingsApi.md#settingsdiscoverslideridput) | **Put** /settings/discover/{sliderId} | Update a single slider
*SettingsApi* | [**SettingsInitializePost**](docs/SettingsApi.md#settingsinitializepost) | **Post** /settings/initialize | Initialize application
*SettingsApi* | [**SettingsJobsGet**](docs/SettingsApi.md#settingsjobsget) | **Get** /settings/jobs | Get scheduled jobs
*SettingsApi* | [**SettingsJobsJobIdCancelPost**](docs/SettingsApi.md#settingsjobsjobidcancelpost) | **Post** /settings/jobs/{jobId}/cancel | Cancel a specific job
*SettingsApi* | [**SettingsJobsJobIdRunPost**](docs/SettingsApi.md#settingsjobsjobidrunpost) | **Post** /settings/jobs/{jobId}/run | Invoke a specific job
*SettingsApi* | [**SettingsJobsJobIdSchedulePost**](docs/SettingsApi.md#settingsjobsjobidschedulepost) | **Post** /settings/jobs/{jobId}/schedule | Modify job schedule
*SettingsApi* | [**SettingsLogsGet**](docs/SettingsApi.md#settingslogsget) | **Get** /settings/logs | Returns logs
*SettingsApi* | [**SettingsMainGet**](docs/SettingsApi.md#settingsmainget) | **Get** /settings/main | Get main settings
*SettingsApi* | [**SettingsMainPost**](docs/SettingsApi.md#settingsmainpost) | **Post** /settings/main | Update main settings
*SettingsApi* | [**SettingsMainRegeneratePost**](docs/SettingsApi.md#settingsmainregeneratepost) | **Post** /settings/main/regenerate | Get main settings with newly-generated API key
*SettingsApi* | [**SettingsNotificationsDiscordGet**](docs/SettingsApi.md#settingsnotificationsdiscordget) | **Get** /settings/notifications/discord | Get Discord notification settings
*SettingsApi* | [**SettingsNotificationsDiscordPost**](docs/SettingsApi.md#settingsnotificationsdiscordpost) | **Post** /settings/notifications/discord | Update Discord notification settings
*SettingsApi* | [**SettingsNotificationsDiscordTestPost**](docs/SettingsApi.md#settingsnotificationsdiscordtestpost) | **Post** /settings/notifications/discord/test | Test Discord settings
*SettingsApi* | [**SettingsNotificationsEmailGet**](docs/SettingsApi.md#settingsnotificationsemailget) | **Get** /settings/notifications/email | Get email notification settings
*SettingsApi* | [**SettingsNotificationsEmailPost**](docs/SettingsApi.md#settingsnotificationsemailpost) | **Post** /settings/notifications/email | Update email notification settings
*SettingsApi* | [**SettingsNotificationsEmailTestPost**](docs/SettingsApi.md#settingsnotificationsemailtestpost) | **Post** /settings/notifications/email/test | Test email settings
*SettingsApi* | [**SettingsNotificationsGotifyGet**](docs/SettingsApi.md#settingsnotificationsgotifyget) | **Get** /settings/notifications/gotify | Get Gotify notification settings
*SettingsApi* | [**SettingsNotificationsGotifyPost**](docs/SettingsApi.md#settingsnotificationsgotifypost) | **Post** /settings/notifications/gotify | Update Gotify notification settings
*SettingsApi* | [**SettingsNotificationsGotifyTestPost**](docs/SettingsApi.md#settingsnotificationsgotifytestpost) | **Post** /settings/notifications/gotify/test | Test Gotify settings
*SettingsApi* | [**SettingsNotificationsLunaseaGet**](docs/SettingsApi.md#settingsnotificationslunaseaget) | **Get** /settings/notifications/lunasea | Get LunaSea notification settings
*SettingsApi* | [**SettingsNotificationsLunaseaPost**](docs/SettingsApi.md#settingsnotificationslunaseapost) | **Post** /settings/notifications/lunasea | Update LunaSea notification settings
*SettingsApi* | [**SettingsNotificationsLunaseaTestPost**](docs/SettingsApi.md#settingsnotificationslunaseatestpost) | **Post** /settings/notifications/lunasea/test | Test LunaSea settings
*SettingsApi* | [**SettingsNotificationsPushbulletGet**](docs/SettingsApi.md#settingsnotificationspushbulletget) | **Get** /settings/notifications/pushbullet | Get Pushbullet notification settings
*SettingsApi* | [**SettingsNotificationsPushbulletPost**](docs/SettingsApi.md#settingsnotificationspushbulletpost) | **Post** /settings/notifications/pushbullet | Update Pushbullet notification settings
*SettingsApi* | [**SettingsNotificationsPushbulletTestPost**](docs/SettingsApi.md#settingsnotificationspushbullettestpost) | **Post** /settings/notifications/pushbullet/test | Test Pushbullet settings
*SettingsApi* | [**SettingsNotificationsPushoverGet**](docs/SettingsApi.md#settingsnotificationspushoverget) | **Get** /settings/notifications/pushover | Get Pushover notification settings
*SettingsApi* | [**SettingsNotificationsPushoverPost**](docs/SettingsApi.md#settingsnotificationspushoverpost) | **Post** /settings/notifications/pushover | Update Pushover notification settings
*SettingsApi* | [**SettingsNotificationsPushoverSoundsGet**](docs/SettingsApi.md#settingsnotificationspushoversoundsget) | **Get** /settings/notifications/pushover/sounds | Get Pushover sounds
*SettingsApi* | [**SettingsNotificationsPushoverTestPost**](docs/SettingsApi.md#settingsnotificationspushovertestpost) | **Post** /settings/notifications/pushover/test | Test Pushover settings
*SettingsApi* | [**SettingsNotificationsSlackGet**](docs/SettingsApi.md#settingsnotificationsslackget) | **Get** /settings/notifications/slack | Get Slack notification settings
*SettingsApi* | [**SettingsNotificationsSlackPost**](docs/SettingsApi.md#settingsnotificationsslackpost) | **Post** /settings/notifications/slack | Update Slack notification settings
*SettingsApi* | [**SettingsNotificationsSlackTestPost**](docs/SettingsApi.md#settingsnotificationsslacktestpost) | **Post** /settings/notifications/slack/test | Test Slack settings
*SettingsApi* | [**SettingsNotificationsTelegramGet**](docs/SettingsApi.md#settingsnotificationstelegramget) | **Get** /settings/notifications/telegram | Get Telegram notification settings
*SettingsApi* | [**SettingsNotificationsTelegramPost**](docs/SettingsApi.md#settingsnotificationstelegrampost) | **Post** /settings/notifications/telegram | Update Telegram notification settings
*SettingsApi* | [**SettingsNotificationsTelegramTestPost**](docs/SettingsApi.md#settingsnotificationstelegramtestpost) | **Post** /settings/notifications/telegram/test | Test Telegram settings
*SettingsApi* | [**SettingsNotificationsWebhookGet**](docs/SettingsApi.md#settingsnotificationswebhookget) | **Get** /settings/notifications/webhook | Get webhook notification settings
*SettingsApi* | [**SettingsNotificationsWebhookPost**](docs/SettingsApi.md#settingsnotificationswebhookpost) | **Post** /settings/notifications/webhook | Update webhook notification settings
*SettingsApi* | [**SettingsNotificationsWebhookTestPost**](docs/SettingsApi.md#settingsnotificationswebhooktestpost) | **Post** /settings/notifications/webhook/test | Test webhook settings
*SettingsApi* | [**SettingsNotificationsWebpushGet**](docs/SettingsApi.md#settingsnotificationswebpushget) | **Get** /settings/notifications/webpush | Get Web Push notification settings
*SettingsApi* | [**SettingsNotificationsWebpushPost**](docs/SettingsApi.md#settingsnotificationswebpushpost) | **Post** /settings/notifications/webpush | Update Web Push notification settings
*SettingsApi* | [**SettingsNotificationsWebpushTestPost**](docs/SettingsApi.md#settingsnotificationswebpushtestpost) | **Post** /settings/notifications/webpush/test | Test Web Push settings
*SettingsApi* | [**SettingsPlexDevicesServersGet**](docs/SettingsApi.md#settingsplexdevicesserversget) | **Get** /settings/plex/devices/servers | Gets the user&#x27;s available Plex servers
*SettingsApi* | [**SettingsPlexGet**](docs/SettingsApi.md#settingsplexget) | **Get** /settings/plex | Get Plex settings
*SettingsApi* | [**SettingsPlexLibraryGet**](docs/SettingsApi.md#settingsplexlibraryget) | **Get** /settings/plex/library | Get Plex libraries
*SettingsApi* | [**SettingsPlexPost**](docs/SettingsApi.md#settingsplexpost) | **Post** /settings/plex | Update Plex settings
*SettingsApi* | [**SettingsPlexSyncGet**](docs/SettingsApi.md#settingsplexsyncget) | **Get** /settings/plex/sync | Get status of full Plex library scan
*SettingsApi* | [**SettingsPlexSyncPost**](docs/SettingsApi.md#settingsplexsyncpost) | **Post** /settings/plex/sync | Start full Plex library scan
*SettingsApi* | [**SettingsPlexUsersGet**](docs/SettingsApi.md#settingsplexusersget) | **Get** /settings/plex/users | Get Plex users
*SettingsApi* | [**SettingsPublicGet**](docs/SettingsApi.md#settingspublicget) | **Get** /settings/public | Get public settings
*SettingsApi* | [**SettingsRadarrGet**](docs/SettingsApi.md#settingsradarrget) | **Get** /settings/radarr | Get Radarr settings
*SettingsApi* | [**SettingsRadarrPost**](docs/SettingsApi.md#settingsradarrpost) | **Post** /settings/radarr | Create Radarr instance
*SettingsApi* | [**SettingsRadarrRadarrIdDelete**](docs/SettingsApi.md#settingsradarrradarriddelete) | **Delete** /settings/radarr/{radarrId} | Delete Radarr instance
*SettingsApi* | [**SettingsRadarrRadarrIdProfilesGet**](docs/SettingsApi.md#settingsradarrradarridprofilesget) | **Get** /settings/radarr/{radarrId}/profiles | Get available Radarr profiles
*SettingsApi* | [**SettingsRadarrRadarrIdPut**](docs/SettingsApi.md#settingsradarrradarridput) | **Put** /settings/radarr/{radarrId} | Update Radarr instance
*SettingsApi* | [**SettingsRadarrTestPost**](docs/SettingsApi.md#settingsradarrtestpost) | **Post** /settings/radarr/test | Test Radarr configuration
*SettingsApi* | [**SettingsSonarrGet**](docs/SettingsApi.md#settingssonarrget) | **Get** /settings/sonarr | Get Sonarr settings
*SettingsApi* | [**SettingsSonarrPost**](docs/SettingsApi.md#settingssonarrpost) | **Post** /settings/sonarr | Create Sonarr instance
*SettingsApi* | [**SettingsSonarrSonarrIdDelete**](docs/SettingsApi.md#settingssonarrsonarriddelete) | **Delete** /settings/sonarr/{sonarrId} | Delete Sonarr instance
*SettingsApi* | [**SettingsSonarrSonarrIdPut**](docs/SettingsApi.md#settingssonarrsonarridput) | **Put** /settings/sonarr/{sonarrId} | Update Sonarr instance
*SettingsApi* | [**SettingsSonarrTestPost**](docs/SettingsApi.md#settingssonarrtestpost) | **Post** /settings/sonarr/test | Test Sonarr configuration
*SettingsApi* | [**SettingsTautulliGet**](docs/SettingsApi.md#settingstautulliget) | **Get** /settings/tautulli | Get Tautulli settings
*SettingsApi* | [**SettingsTautulliPost**](docs/SettingsApi.md#settingstautullipost) | **Post** /settings/tautulli | Update Tautulli settings
*TmdbApi* | [**BackdropsGet**](docs/TmdbApi.md#backdropsget) | **Get** /backdrops | Get backdrops of trending items
*TmdbApi* | [**GenresMovieGet**](docs/TmdbApi.md#genresmovieget) | **Get** /genres/movie | Get list of official TMDB movie genres
*TmdbApi* | [**GenresTvGet**](docs/TmdbApi.md#genrestvget) | **Get** /genres/tv | Get list of official TMDB movie genres
*TmdbApi* | [**LanguagesGet**](docs/TmdbApi.md#languagesget) | **Get** /languages | Languages supported by TMDB
*TmdbApi* | [**NetworkNetworkIdGet**](docs/TmdbApi.md#networknetworkidget) | **Get** /network/{networkId} | Get TV network details
*TmdbApi* | [**RegionsGet**](docs/TmdbApi.md#regionsget) | **Get** /regions | Regions supported by TMDB
*TmdbApi* | [**StudioStudioIdGet**](docs/TmdbApi.md#studiostudioidget) | **Get** /studio/{studioId} | Get movie studio details
*TvApi* | [**TvTvIdGet**](docs/TvApi.md#tvtvidget) | **Get** /tv/{tvId} | Get TV details
*TvApi* | [**TvTvIdRatingsGet**](docs/TvApi.md#tvtvidratingsget) | **Get** /tv/{tvId}/ratings | Get TV ratings
*TvApi* | [**TvTvIdRecommendationsGet**](docs/TvApi.md#tvtvidrecommendationsget) | **Get** /tv/{tvId}/recommendations | Get recommended TV series
*TvApi* | [**TvTvIdSeasonSeasonIdGet**](docs/TvApi.md#tvtvidseasonseasonidget) | **Get** /tv/{tvId}/season/{seasonId} | Get season details and episode list
*TvApi* | [**TvTvIdSimilarGet**](docs/TvApi.md#tvtvidsimilarget) | **Get** /tv/{tvId}/similar | Get similar TV series
*UsersApi* | [**AuthMeGet**](docs/UsersApi.md#authmeget) | **Get** /auth/me | Get logged-in user
*UsersApi* | [**AuthResetPasswordGuidPost**](docs/UsersApi.md#authresetpasswordguidpost) | **Post** /auth/reset-password/{guid} | Reset the password for a user
*UsersApi* | [**AuthResetPasswordPost**](docs/UsersApi.md#authresetpasswordpost) | **Post** /auth/reset-password | Send a reset password email
*UsersApi* | [**SettingsPlexUsersGet**](docs/UsersApi.md#settingsplexusersget) | **Get** /settings/plex/users | Get Plex users
*UsersApi* | [**UserGet**](docs/UsersApi.md#userget) | **Get** /user | Get all users
*UsersApi* | [**UserImportFromPlexPost**](docs/UsersApi.md#userimportfromplexpost) | **Post** /user/import-from-plex | Import all users from Plex
*UsersApi* | [**UserPost**](docs/UsersApi.md#userpost) | **Post** /user | Create new user
*UsersApi* | [**UserPut**](docs/UsersApi.md#userput) | **Put** /user | Update batch of users
*UsersApi* | [**UserRegisterPushSubscriptionPost**](docs/UsersApi.md#userregisterpushsubscriptionpost) | **Post** /user/registerPushSubscription | Register a web push /user/registerPushSubscription
*UsersApi* | [**UserUserIdDelete**](docs/UsersApi.md#useruseriddelete) | **Delete** /user/{userId} | Delete user by ID
*UsersApi* | [**UserUserIdGet**](docs/UsersApi.md#useruseridget) | **Get** /user/{userId} | Get user by ID
*UsersApi* | [**UserUserIdPut**](docs/UsersApi.md#useruseridput) | **Put** /user/{userId} | Update a user by user ID
*UsersApi* | [**UserUserIdQuotaGet**](docs/UsersApi.md#useruseridquotaget) | **Get** /user/{userId}/quota | Get quotas for a specific user
*UsersApi* | [**UserUserIdRequestsGet**](docs/UsersApi.md#useruseridrequestsget) | **Get** /user/{userId}/requests | Get requests for a specific user
*UsersApi* | [**UserUserIdSettingsMainGet**](docs/UsersApi.md#useruseridsettingsmainget) | **Get** /user/{userId}/settings/main | Get general settings for a user
*UsersApi* | [**UserUserIdSettingsMainPost**](docs/UsersApi.md#useruseridsettingsmainpost) | **Post** /user/{userId}/settings/main | Update general settings for a user
*UsersApi* | [**UserUserIdSettingsNotificationsGet**](docs/UsersApi.md#useruseridsettingsnotificationsget) | **Get** /user/{userId}/settings/notifications | Get notification settings for a user
*UsersApi* | [**UserUserIdSettingsNotificationsPost**](docs/UsersApi.md#useruseridsettingsnotificationspost) | **Post** /user/{userId}/settings/notifications | Update notification settings for a user
*UsersApi* | [**UserUserIdSettingsPasswordGet**](docs/UsersApi.md#useruseridsettingspasswordget) | **Get** /user/{userId}/settings/password | Get password page informatiom
*UsersApi* | [**UserUserIdSettingsPasswordPost**](docs/UsersApi.md#useruseridsettingspasswordpost) | **Post** /user/{userId}/settings/password | Update password for a user
*UsersApi* | [**UserUserIdSettingsPermissionsGet**](docs/UsersApi.md#useruseridsettingspermissionsget) | **Get** /user/{userId}/settings/permissions | Get permission settings for a user
*UsersApi* | [**UserUserIdSettingsPermissionsPost**](docs/UsersApi.md#useruseridsettingspermissionspost) | **Post** /user/{userId}/settings/permissions | Update permission settings for a user
*UsersApi* | [**UserUserIdWatchDataGet**](docs/UsersApi.md#useruseridwatchdataget) | **Get** /user/{userId}/watch_data | Get watch data
*UsersApi* | [**UserUserIdWatchlistGet**](docs/UsersApi.md#useruseridwatchlistget) | **Get** /user/{userId}/watchlist | Get the Plex watchlist for a specific user

## Documentation For Models

 - [AnyOfMediaRequestModifiedBy](docs/AnyOfMediaRequestModifiedBy.md)
 - [AnyOfinlineResponse20018ResultsItems](docs/AnyOfinlineResponse20018ResultsItems.md)
 - [AuthLocalBody](docs/AuthLocalBody.md)
 - [AuthPlexBody](docs/AuthPlexBody.md)
 - [AuthResetpasswordBody](docs/AuthResetpasswordBody.md)
 - [Cast](docs/Cast.md)
 - [Collection](docs/Collection.md)
 - [Company](docs/Company.md)
 - [CreditCast](docs/CreditCast.md)
 - [CreditCrew](docs/CreditCrew.md)
 - [Crew](docs/Crew.md)
 - [DiscordSettings](docs/DiscordSettings.md)
 - [DiscordSettingsOptions](docs/DiscordSettingsOptions.md)
 - [DiscoverAddBody](docs/DiscoverAddBody.md)
 - [DiscoverSlider](docs/DiscoverSlider.md)
 - [DiscoverSliderIdBody](docs/DiscoverSliderIdBody.md)
 - [Episode](docs/Episode.md)
 - [ExternalIds](docs/ExternalIds.md)
 - [Genre](docs/Genre.md)
 - [GotifySettings](docs/GotifySettings.md)
 - [GotifySettingsOptions](docs/GotifySettingsOptions.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse20012Movie](docs/InlineResponse20012Movie.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse20013Results](docs/InlineResponse20013Results.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse20019](docs/InlineResponse20019.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse20020](docs/InlineResponse20020.md)
 - [InlineResponse20021](docs/InlineResponse20021.md)
 - [InlineResponse20022](docs/InlineResponse20022.md)
 - [InlineResponse20023](docs/InlineResponse20023.md)
 - [InlineResponse20024](docs/InlineResponse20024.md)
 - [InlineResponse20025](docs/InlineResponse20025.md)
 - [InlineResponse20026](docs/InlineResponse20026.md)
 - [InlineResponse20027](docs/InlineResponse20027.md)
 - [InlineResponse20028](docs/InlineResponse20028.md)
 - [InlineResponse20029](docs/InlineResponse20029.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse20030](docs/InlineResponse20030.md)
 - [InlineResponse20031](docs/InlineResponse20031.md)
 - [InlineResponse20032](docs/InlineResponse20032.md)
 - [InlineResponse20032Imdb](docs/InlineResponse20032Imdb.md)
 - [InlineResponse20033](docs/InlineResponse20033.md)
 - [InlineResponse20034](docs/InlineResponse20034.md)
 - [InlineResponse20035](docs/InlineResponse20035.md)
 - [InlineResponse20036](docs/InlineResponse20036.md)
 - [InlineResponse20036Data](docs/InlineResponse20036Data.md)
 - [InlineResponse20037](docs/InlineResponse20037.md)
 - [InlineResponse20038](docs/InlineResponse20038.md)
 - [InlineResponse20039](docs/InlineResponse20039.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse20040](docs/InlineResponse20040.md)
 - [InlineResponse20041](docs/InlineResponse20041.md)
 - [InlineResponse20042](docs/InlineResponse20042.md)
 - [InlineResponse20043](docs/InlineResponse20043.md)
 - [InlineResponse20044](docs/InlineResponse20044.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2005ApiCaches](docs/InlineResponse2005ApiCaches.md)
 - [InlineResponse2005ImageCache](docs/InlineResponse2005ImageCache.md)
 - [InlineResponse2005ImageCacheTmdb](docs/InlineResponse2005ImageCacheTmdb.md)
 - [InlineResponse2005Stats](docs/InlineResponse2005Stats.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [Issue](docs/Issue.md)
 - [IssueBody](docs/IssueBody.md)
 - [IssueComment](docs/IssueComment.md)
 - [IssueCommentCommentIdBody](docs/IssueCommentCommentIdBody.md)
 - [IssueIdCommentBody](docs/IssueIdCommentBody.md)
 - [Job](docs/Job.md)
 - [JobIdScheduleBody](docs/JobIdScheduleBody.md)
 - [Keyword](docs/Keyword.md)
 - [LunaSeaSettings](docs/LunaSeaSettings.md)
 - [LunaSeaSettingsOptions](docs/LunaSeaSettingsOptions.md)
 - [MainSettings](docs/MainSettings.md)
 - [MediaIdStatusBody](docs/MediaIdStatusBody.md)
 - [MediaInfo](docs/MediaInfo.md)
 - [MediaRequest](docs/MediaRequest.md)
 - [MovieDetails](docs/MovieDetails.md)
 - [MovieDetailsCollection](docs/MovieDetailsCollection.md)
 - [MovieDetailsCredits](docs/MovieDetailsCredits.md)
 - [MovieDetailsProductionCountries](docs/MovieDetailsProductionCountries.md)
 - [MovieDetailsReleases](docs/MovieDetailsReleases.md)
 - [MovieDetailsReleasesReleaseDates](docs/MovieDetailsReleasesReleaseDates.md)
 - [MovieDetailsReleasesResults](docs/MovieDetailsReleasesResults.md)
 - [MovieResult](docs/MovieResult.md)
 - [Network](docs/Network.md)
 - [NotificationAgentTypes](docs/NotificationAgentTypes.md)
 - [NotificationEmailSettings](docs/NotificationEmailSettings.md)
 - [NotificationEmailSettingsOptions](docs/NotificationEmailSettingsOptions.md)
 - [OneOfPersonResultKnownForItems](docs/OneOfPersonResultKnownForItems.md)
 - [OneOfrequestBodySeasons](docs/OneOfrequestBodySeasons.md)
 - [PageInfo](docs/PageInfo.md)
 - [PersonDetails](docs/PersonDetails.md)
 - [PersonResult](docs/PersonResult.md)
 - [PlexConnection](docs/PlexConnection.md)
 - [PlexDevice](docs/PlexDevice.md)
 - [PlexLibrary](docs/PlexLibrary.md)
 - [PlexSettings](docs/PlexSettings.md)
 - [PlexSyncBody](docs/PlexSyncBody.md)
 - [ProductionCompany](docs/ProductionCompany.md)
 - [PublicSettings](docs/PublicSettings.md)
 - [PushbulletSettings](docs/PushbulletSettings.md)
 - [PushbulletSettingsOptions](docs/PushbulletSettingsOptions.md)
 - [PushoverSettings](docs/PushoverSettings.md)
 - [PushoverSettingsOptions](docs/PushoverSettingsOptions.md)
 - [RadarrSettings](docs/RadarrSettings.md)
 - [RadarrTestBody](docs/RadarrTestBody.md)
 - [RelatedVideo](docs/RelatedVideo.md)
 - [RequestBody](docs/RequestBody.md)
 - [RequestRequestIdBody](docs/RequestRequestIdBody.md)
 - [ResetpasswordGuidBody](docs/ResetpasswordGuidBody.md)
 - [Season](docs/Season.md)
 - [ServarrTag](docs/ServarrTag.md)
 - [ServiceProfile](docs/ServiceProfile.md)
 - [SettingsMainBody](docs/SettingsMainBody.md)
 - [SettingsPasswordBody](docs/SettingsPasswordBody.md)
 - [SettingsPermissionsBody](docs/SettingsPermissionsBody.md)
 - [SlackSettings](docs/SlackSettings.md)
 - [SlackSettingsOptions](docs/SlackSettingsOptions.md)
 - [SonarrSeries](docs/SonarrSeries.md)
 - [SonarrSeriesAddOptions](docs/SonarrSeriesAddOptions.md)
 - [SonarrSeriesImages](docs/SonarrSeriesImages.md)
 - [SonarrSeriesRatings](docs/SonarrSeriesRatings.md)
 - [SonarrSeriesSeasons](docs/SonarrSeriesSeasons.md)
 - [SonarrSettings](docs/SonarrSettings.md)
 - [SonarrTestBody](docs/SonarrTestBody.md)
 - [SpokenLanguage](docs/SpokenLanguage.md)
 - [TautulliSettings](docs/TautulliSettings.md)
 - [TelegramSettings](docs/TelegramSettings.md)
 - [TelegramSettingsOptions](docs/TelegramSettingsOptions.md)
 - [TvDetails](docs/TvDetails.md)
 - [TvDetailsContentRatings](docs/TvDetailsContentRatings.md)
 - [TvDetailsContentRatingsResults](docs/TvDetailsContentRatingsResults.md)
 - [TvDetailsCreatedBy](docs/TvDetailsCreatedBy.md)
 - [TvResult](docs/TvResult.md)
 - [User](docs/User.md)
 - [UserBody](docs/UserBody.md)
 - [UserBody1](docs/UserBody1.md)
 - [UserImportfromplexBody](docs/UserImportfromplexBody.md)
 - [UserRegisterPushSubscriptionBody](docs/UserRegisterPushSubscriptionBody.md)
 - [UserSettings](docs/UserSettings.md)
 - [UserSettingsNotifications](docs/UserSettingsNotifications.md)
 - [WatchProviderDetails](docs/WatchProviderDetails.md)
 - [WatchProviderRegion](docs/WatchProviderRegion.md)
 - [WatchProvidersInner](docs/WatchProvidersInner.md)
 - [WebPushSettings](docs/WebPushSettings.md)
 - [WebhookSettings](docs/WebhookSettings.md)
 - [WebhookSettingsOptions](docs/WebhookSettingsOptions.md)

## Documentation For Authorization

## apiKey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```
## cookieAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

