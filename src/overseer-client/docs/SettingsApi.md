# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsAboutGet**](SettingsApi.md#SettingsAboutGet) | **Get** /settings/about | Get server stats
[**SettingsCacheCacheIdFlushPost**](SettingsApi.md#SettingsCacheCacheIdFlushPost) | **Post** /settings/cache/{cacheId}/flush | Flush a specific cache
[**SettingsCacheGet**](SettingsApi.md#SettingsCacheGet) | **Get** /settings/cache | Get a list of active caches
[**SettingsDiscoverAddPost**](SettingsApi.md#SettingsDiscoverAddPost) | **Post** /settings/discover/add | Add a new slider
[**SettingsDiscoverGet**](SettingsApi.md#SettingsDiscoverGet) | **Get** /settings/discover | Get all discover sliders
[**SettingsDiscoverPost**](SettingsApi.md#SettingsDiscoverPost) | **Post** /settings/discover | Batch update all sliders.
[**SettingsDiscoverResetGet**](SettingsApi.md#SettingsDiscoverResetGet) | **Get** /settings/discover/reset | Reset all discover sliders
[**SettingsDiscoverSliderIdDelete**](SettingsApi.md#SettingsDiscoverSliderIdDelete) | **Delete** /settings/discover/{sliderId} | Delete slider by ID
[**SettingsDiscoverSliderIdPut**](SettingsApi.md#SettingsDiscoverSliderIdPut) | **Put** /settings/discover/{sliderId} | Update a single slider
[**SettingsInitializePost**](SettingsApi.md#SettingsInitializePost) | **Post** /settings/initialize | Initialize application
[**SettingsJobsGet**](SettingsApi.md#SettingsJobsGet) | **Get** /settings/jobs | Get scheduled jobs
[**SettingsJobsJobIdCancelPost**](SettingsApi.md#SettingsJobsJobIdCancelPost) | **Post** /settings/jobs/{jobId}/cancel | Cancel a specific job
[**SettingsJobsJobIdRunPost**](SettingsApi.md#SettingsJobsJobIdRunPost) | **Post** /settings/jobs/{jobId}/run | Invoke a specific job
[**SettingsJobsJobIdSchedulePost**](SettingsApi.md#SettingsJobsJobIdSchedulePost) | **Post** /settings/jobs/{jobId}/schedule | Modify job schedule
[**SettingsLogsGet**](SettingsApi.md#SettingsLogsGet) | **Get** /settings/logs | Returns logs
[**SettingsMainGet**](SettingsApi.md#SettingsMainGet) | **Get** /settings/main | Get main settings
[**SettingsMainPost**](SettingsApi.md#SettingsMainPost) | **Post** /settings/main | Update main settings
[**SettingsMainRegeneratePost**](SettingsApi.md#SettingsMainRegeneratePost) | **Post** /settings/main/regenerate | Get main settings with newly-generated API key
[**SettingsNotificationsDiscordGet**](SettingsApi.md#SettingsNotificationsDiscordGet) | **Get** /settings/notifications/discord | Get Discord notification settings
[**SettingsNotificationsDiscordPost**](SettingsApi.md#SettingsNotificationsDiscordPost) | **Post** /settings/notifications/discord | Update Discord notification settings
[**SettingsNotificationsDiscordTestPost**](SettingsApi.md#SettingsNotificationsDiscordTestPost) | **Post** /settings/notifications/discord/test | Test Discord settings
[**SettingsNotificationsEmailGet**](SettingsApi.md#SettingsNotificationsEmailGet) | **Get** /settings/notifications/email | Get email notification settings
[**SettingsNotificationsEmailPost**](SettingsApi.md#SettingsNotificationsEmailPost) | **Post** /settings/notifications/email | Update email notification settings
[**SettingsNotificationsEmailTestPost**](SettingsApi.md#SettingsNotificationsEmailTestPost) | **Post** /settings/notifications/email/test | Test email settings
[**SettingsNotificationsGotifyGet**](SettingsApi.md#SettingsNotificationsGotifyGet) | **Get** /settings/notifications/gotify | Get Gotify notification settings
[**SettingsNotificationsGotifyPost**](SettingsApi.md#SettingsNotificationsGotifyPost) | **Post** /settings/notifications/gotify | Update Gotify notification settings
[**SettingsNotificationsGotifyTestPost**](SettingsApi.md#SettingsNotificationsGotifyTestPost) | **Post** /settings/notifications/gotify/test | Test Gotify settings
[**SettingsNotificationsLunaseaGet**](SettingsApi.md#SettingsNotificationsLunaseaGet) | **Get** /settings/notifications/lunasea | Get LunaSea notification settings
[**SettingsNotificationsLunaseaPost**](SettingsApi.md#SettingsNotificationsLunaseaPost) | **Post** /settings/notifications/lunasea | Update LunaSea notification settings
[**SettingsNotificationsLunaseaTestPost**](SettingsApi.md#SettingsNotificationsLunaseaTestPost) | **Post** /settings/notifications/lunasea/test | Test LunaSea settings
[**SettingsNotificationsPushbulletGet**](SettingsApi.md#SettingsNotificationsPushbulletGet) | **Get** /settings/notifications/pushbullet | Get Pushbullet notification settings
[**SettingsNotificationsPushbulletPost**](SettingsApi.md#SettingsNotificationsPushbulletPost) | **Post** /settings/notifications/pushbullet | Update Pushbullet notification settings
[**SettingsNotificationsPushbulletTestPost**](SettingsApi.md#SettingsNotificationsPushbulletTestPost) | **Post** /settings/notifications/pushbullet/test | Test Pushbullet settings
[**SettingsNotificationsPushoverGet**](SettingsApi.md#SettingsNotificationsPushoverGet) | **Get** /settings/notifications/pushover | Get Pushover notification settings
[**SettingsNotificationsPushoverPost**](SettingsApi.md#SettingsNotificationsPushoverPost) | **Post** /settings/notifications/pushover | Update Pushover notification settings
[**SettingsNotificationsPushoverSoundsGet**](SettingsApi.md#SettingsNotificationsPushoverSoundsGet) | **Get** /settings/notifications/pushover/sounds | Get Pushover sounds
[**SettingsNotificationsPushoverTestPost**](SettingsApi.md#SettingsNotificationsPushoverTestPost) | **Post** /settings/notifications/pushover/test | Test Pushover settings
[**SettingsNotificationsSlackGet**](SettingsApi.md#SettingsNotificationsSlackGet) | **Get** /settings/notifications/slack | Get Slack notification settings
[**SettingsNotificationsSlackPost**](SettingsApi.md#SettingsNotificationsSlackPost) | **Post** /settings/notifications/slack | Update Slack notification settings
[**SettingsNotificationsSlackTestPost**](SettingsApi.md#SettingsNotificationsSlackTestPost) | **Post** /settings/notifications/slack/test | Test Slack settings
[**SettingsNotificationsTelegramGet**](SettingsApi.md#SettingsNotificationsTelegramGet) | **Get** /settings/notifications/telegram | Get Telegram notification settings
[**SettingsNotificationsTelegramPost**](SettingsApi.md#SettingsNotificationsTelegramPost) | **Post** /settings/notifications/telegram | Update Telegram notification settings
[**SettingsNotificationsTelegramTestPost**](SettingsApi.md#SettingsNotificationsTelegramTestPost) | **Post** /settings/notifications/telegram/test | Test Telegram settings
[**SettingsNotificationsWebhookGet**](SettingsApi.md#SettingsNotificationsWebhookGet) | **Get** /settings/notifications/webhook | Get webhook notification settings
[**SettingsNotificationsWebhookPost**](SettingsApi.md#SettingsNotificationsWebhookPost) | **Post** /settings/notifications/webhook | Update webhook notification settings
[**SettingsNotificationsWebhookTestPost**](SettingsApi.md#SettingsNotificationsWebhookTestPost) | **Post** /settings/notifications/webhook/test | Test webhook settings
[**SettingsNotificationsWebpushGet**](SettingsApi.md#SettingsNotificationsWebpushGet) | **Get** /settings/notifications/webpush | Get Web Push notification settings
[**SettingsNotificationsWebpushPost**](SettingsApi.md#SettingsNotificationsWebpushPost) | **Post** /settings/notifications/webpush | Update Web Push notification settings
[**SettingsNotificationsWebpushTestPost**](SettingsApi.md#SettingsNotificationsWebpushTestPost) | **Post** /settings/notifications/webpush/test | Test Web Push settings
[**SettingsPlexDevicesServersGet**](SettingsApi.md#SettingsPlexDevicesServersGet) | **Get** /settings/plex/devices/servers | Gets the user&#x27;s available Plex servers
[**SettingsPlexGet**](SettingsApi.md#SettingsPlexGet) | **Get** /settings/plex | Get Plex settings
[**SettingsPlexLibraryGet**](SettingsApi.md#SettingsPlexLibraryGet) | **Get** /settings/plex/library | Get Plex libraries
[**SettingsPlexPost**](SettingsApi.md#SettingsPlexPost) | **Post** /settings/plex | Update Plex settings
[**SettingsPlexSyncGet**](SettingsApi.md#SettingsPlexSyncGet) | **Get** /settings/plex/sync | Get status of full Plex library scan
[**SettingsPlexSyncPost**](SettingsApi.md#SettingsPlexSyncPost) | **Post** /settings/plex/sync | Start full Plex library scan
[**SettingsPlexUsersGet**](SettingsApi.md#SettingsPlexUsersGet) | **Get** /settings/plex/users | Get Plex users
[**SettingsPublicGet**](SettingsApi.md#SettingsPublicGet) | **Get** /settings/public | Get public settings
[**SettingsRadarrGet**](SettingsApi.md#SettingsRadarrGet) | **Get** /settings/radarr | Get Radarr settings
[**SettingsRadarrPost**](SettingsApi.md#SettingsRadarrPost) | **Post** /settings/radarr | Create Radarr instance
[**SettingsRadarrRadarrIdDelete**](SettingsApi.md#SettingsRadarrRadarrIdDelete) | **Delete** /settings/radarr/{radarrId} | Delete Radarr instance
[**SettingsRadarrRadarrIdProfilesGet**](SettingsApi.md#SettingsRadarrRadarrIdProfilesGet) | **Get** /settings/radarr/{radarrId}/profiles | Get available Radarr profiles
[**SettingsRadarrRadarrIdPut**](SettingsApi.md#SettingsRadarrRadarrIdPut) | **Put** /settings/radarr/{radarrId} | Update Radarr instance
[**SettingsRadarrTestPost**](SettingsApi.md#SettingsRadarrTestPost) | **Post** /settings/radarr/test | Test Radarr configuration
[**SettingsSonarrGet**](SettingsApi.md#SettingsSonarrGet) | **Get** /settings/sonarr | Get Sonarr settings
[**SettingsSonarrPost**](SettingsApi.md#SettingsSonarrPost) | **Post** /settings/sonarr | Create Sonarr instance
[**SettingsSonarrSonarrIdDelete**](SettingsApi.md#SettingsSonarrSonarrIdDelete) | **Delete** /settings/sonarr/{sonarrId} | Delete Sonarr instance
[**SettingsSonarrSonarrIdPut**](SettingsApi.md#SettingsSonarrSonarrIdPut) | **Put** /settings/sonarr/{sonarrId} | Update Sonarr instance
[**SettingsSonarrTestPost**](SettingsApi.md#SettingsSonarrTestPost) | **Post** /settings/sonarr/test | Test Sonarr configuration
[**SettingsTautulliGet**](SettingsApi.md#SettingsTautulliGet) | **Get** /settings/tautulli | Get Tautulli settings
[**SettingsTautulliPost**](SettingsApi.md#SettingsTautulliPost) | **Post** /settings/tautulli | Update Tautulli settings

# **SettingsAboutGet**
> InlineResponse2008 SettingsAboutGet(ctx, )
Get server stats

Returns current server stats in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsCacheCacheIdFlushPost**
> SettingsCacheCacheIdFlushPost(ctx, cacheId)
Flush a specific cache

Flushes all data from the cache ID provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cacheId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsCacheGet**
> InlineResponse2005 SettingsCacheGet(ctx, )
Get a list of active caches

Retrieves a list of all active caches and their current stats.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsDiscoverAddPost**
> DiscoverSlider SettingsDiscoverAddPost(ctx, body)
Add a new slider

Add a single slider and return the newly created slider. Requires the `ADMIN` permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DiscoverAddBody**](DiscoverAddBody.md)|  | 

### Return type

[**DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsDiscoverGet**
> []DiscoverSlider SettingsDiscoverGet(ctx, )
Get all discover sliders

Returns all discovery sliders. Built-in and custom made.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsDiscoverPost**
> []DiscoverSlider SettingsDiscoverPost(ctx, body)
Batch update all sliders.

Batch update all sliders at once. Should also be used for creation. Will only update sliders provided and will not delete any sliders not present in the request. If a slider is missing a required field, it will be ignored. Requires the `ADMIN` permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]DiscoverSlider**](DiscoverSlider.md)|  | 

### Return type

[**[]DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsDiscoverResetGet**
> SettingsDiscoverResetGet(ctx, )
Reset all discover sliders

Resets all discovery sliders to the default values. Requires the `ADMIN` permission.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsDiscoverSliderIdDelete**
> DiscoverSlider SettingsDiscoverSliderIdDelete(ctx, sliderId)
Delete slider by ID

Deletes the slider with the provided sliderId. Requires the `ADMIN` permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sliderId** | **float64**|  | 

### Return type

[**DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsDiscoverSliderIdPut**
> DiscoverSlider SettingsDiscoverSliderIdPut(ctx, body, sliderId)
Update a single slider

Updates a single slider and return the newly updated slider. Requires the `ADMIN` permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DiscoverSliderIdBody**](DiscoverSliderIdBody.md)|  | 
  **sliderId** | **float64**|  | 

### Return type

[**DiscoverSlider**](DiscoverSlider.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsInitializePost**
> PublicSettings SettingsInitializePost(ctx, )
Initialize application

Sets the app as initialized, allowing the user to navigate to pages other than the setup page.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PublicSettings**](PublicSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsJobsGet**
> []Job SettingsJobsGet(ctx, )
Get scheduled jobs

Returns list of all scheduled jobs and details about their next execution time in a JSON array.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsJobsJobIdCancelPost**
> Job SettingsJobsJobIdCancelPost(ctx, jobId)
Cancel a specific job

Cancels a specific job. Will return the new job status in JSON format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsJobsJobIdRunPost**
> Job SettingsJobsJobIdRunPost(ctx, jobId)
Invoke a specific job

Invokes a specific job to run. Will return the new job status in JSON format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **jobId** | **string**|  | 

### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsJobsJobIdSchedulePost**
> Job SettingsJobsJobIdSchedulePost(ctx, body, jobId)
Modify job schedule

Re-registers the job with the schedule specified. Will return the job in JSON format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JobIdScheduleBody**](JobIdScheduleBody.md)|  | 
  **jobId** | **string**|  | 

### Return type

[**Job**](Job.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsLogsGet**
> []InlineResponse2006 SettingsLogsGet(ctx, optional)
Returns logs

Returns list of all log items and details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiSettingsLogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiSettingsLogsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **optional.Float64**|  | 
 **skip** | **optional.Float64**|  | 
 **filter** | **optional.String**|  | [default to debug]
 **search** | **optional.String**|  | 

### Return type

[**[]InlineResponse2006**](inline_response_200_6.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsMainGet**
> MainSettings SettingsMainGet(ctx, )
Get main settings

Retrieves all main settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsMainPost**
> MainSettings SettingsMainPost(ctx, body)
Update main settings

Updates main settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MainSettings**](MainSettings.md)|  | 

### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsMainRegeneratePost**
> MainSettings SettingsMainRegeneratePost(ctx, )
Get main settings with newly-generated API key

Returns main settings in a JSON object, using the new API key.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**MainSettings**](MainSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsDiscordGet**
> DiscordSettings SettingsNotificationsDiscordGet(ctx, )
Get Discord notification settings

Returns current Discord notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DiscordSettings**](DiscordSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsDiscordPost**
> DiscordSettings SettingsNotificationsDiscordPost(ctx, body)
Update Discord notification settings

Updates Discord notification settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DiscordSettings**](DiscordSettings.md)|  | 

### Return type

[**DiscordSettings**](DiscordSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsDiscordTestPost**
> SettingsNotificationsDiscordTestPost(ctx, body)
Test Discord settings

Sends a test notification to the Discord agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DiscordSettings**](DiscordSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsEmailGet**
> NotificationEmailSettings SettingsNotificationsEmailGet(ctx, )
Get email notification settings

Returns current email notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**NotificationEmailSettings**](NotificationEmailSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsEmailPost**
> NotificationEmailSettings SettingsNotificationsEmailPost(ctx, body)
Update email notification settings

Updates email notification settings with provided values

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationEmailSettings**](NotificationEmailSettings.md)|  | 

### Return type

[**NotificationEmailSettings**](NotificationEmailSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsEmailTestPost**
> SettingsNotificationsEmailTestPost(ctx, body)
Test email settings

Sends a test notification to the email agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationEmailSettings**](NotificationEmailSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsGotifyGet**
> GotifySettings SettingsNotificationsGotifyGet(ctx, )
Get Gotify notification settings

Returns current Gotify notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GotifySettings**](GotifySettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsGotifyPost**
> GotifySettings SettingsNotificationsGotifyPost(ctx, body)
Update Gotify notification settings

Update Gotify notification settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GotifySettings**](GotifySettings.md)|  | 

### Return type

[**GotifySettings**](GotifySettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsGotifyTestPost**
> SettingsNotificationsGotifyTestPost(ctx, body)
Test Gotify settings

Sends a test notification to the Gotify agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GotifySettings**](GotifySettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsLunaseaGet**
> LunaSeaSettings SettingsNotificationsLunaseaGet(ctx, )
Get LunaSea notification settings

Returns current LunaSea notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**LunaSeaSettings**](LunaSeaSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsLunaseaPost**
> LunaSeaSettings SettingsNotificationsLunaseaPost(ctx, body)
Update LunaSea notification settings

Updates LunaSea notification settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LunaSeaSettings**](LunaSeaSettings.md)|  | 

### Return type

[**LunaSeaSettings**](LunaSeaSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsLunaseaTestPost**
> SettingsNotificationsLunaseaTestPost(ctx, body)
Test LunaSea settings

Sends a test notification to the LunaSea agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LunaSeaSettings**](LunaSeaSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsPushbulletGet**
> PushbulletSettings SettingsNotificationsPushbulletGet(ctx, )
Get Pushbullet notification settings

Returns current Pushbullet notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PushbulletSettings**](PushbulletSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsPushbulletPost**
> PushbulletSettings SettingsNotificationsPushbulletPost(ctx, body)
Update Pushbullet notification settings

Update Pushbullet notification settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PushbulletSettings**](PushbulletSettings.md)|  | 

### Return type

[**PushbulletSettings**](PushbulletSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsPushbulletTestPost**
> SettingsNotificationsPushbulletTestPost(ctx, body)
Test Pushbullet settings

Sends a test notification to the Pushbullet agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PushbulletSettings**](PushbulletSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsPushoverGet**
> PushoverSettings SettingsNotificationsPushoverGet(ctx, )
Get Pushover notification settings

Returns current Pushover notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PushoverSettings**](PushoverSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsPushoverPost**
> PushoverSettings SettingsNotificationsPushoverPost(ctx, body)
Update Pushover notification settings

Update Pushover notification settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PushoverSettings**](PushoverSettings.md)|  | 

### Return type

[**PushoverSettings**](PushoverSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsPushoverSoundsGet**
> []InlineResponse2007 SettingsNotificationsPushoverSoundsGet(ctx, token)
Get Pushover sounds

Returns valid Pushover sound options in a JSON array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 

### Return type

[**[]InlineResponse2007**](inline_response_200_7.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsPushoverTestPost**
> SettingsNotificationsPushoverTestPost(ctx, body)
Test Pushover settings

Sends a test notification to the Pushover agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PushoverSettings**](PushoverSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsSlackGet**
> SlackSettings SettingsNotificationsSlackGet(ctx, )
Get Slack notification settings

Returns current Slack notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SlackSettings**](SlackSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsSlackPost**
> SlackSettings SettingsNotificationsSlackPost(ctx, body)
Update Slack notification settings

Updates Slack notification settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SlackSettings**](SlackSettings.md)|  | 

### Return type

[**SlackSettings**](SlackSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsSlackTestPost**
> SettingsNotificationsSlackTestPost(ctx, body)
Test Slack settings

Sends a test notification to the Slack agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SlackSettings**](SlackSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsTelegramGet**
> TelegramSettings SettingsNotificationsTelegramGet(ctx, )
Get Telegram notification settings

Returns current Telegram notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TelegramSettings**](TelegramSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsTelegramPost**
> TelegramSettings SettingsNotificationsTelegramPost(ctx, body)
Update Telegram notification settings

Update Telegram notification settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TelegramSettings**](TelegramSettings.md)|  | 

### Return type

[**TelegramSettings**](TelegramSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsTelegramTestPost**
> SettingsNotificationsTelegramTestPost(ctx, body)
Test Telegram settings

Sends a test notification to the Telegram agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TelegramSettings**](TelegramSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsWebhookGet**
> WebhookSettings SettingsNotificationsWebhookGet(ctx, )
Get webhook notification settings

Returns current webhook notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WebhookSettings**](WebhookSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsWebhookPost**
> WebhookSettings SettingsNotificationsWebhookPost(ctx, body)
Update webhook notification settings

Updates webhook notification settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhookSettings**](WebhookSettings.md)|  | 

### Return type

[**WebhookSettings**](WebhookSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsWebhookTestPost**
> SettingsNotificationsWebhookTestPost(ctx, body)
Test webhook settings

Sends a test notification to the webhook agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebhookSettings**](WebhookSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsWebpushGet**
> WebPushSettings SettingsNotificationsWebpushGet(ctx, )
Get Web Push notification settings

Returns current Web Push notification settings in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**WebPushSettings**](WebPushSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsWebpushPost**
> WebPushSettings SettingsNotificationsWebpushPost(ctx, body)
Update Web Push notification settings

Updates Web Push notification settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebPushSettings**](WebPushSettings.md)|  | 

### Return type

[**WebPushSettings**](WebPushSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsNotificationsWebpushTestPost**
> SettingsNotificationsWebpushTestPost(ctx, body)
Test Web Push settings

Sends a test notification to the Web Push agent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WebPushSettings**](WebPushSettings.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPlexDevicesServersGet**
> []PlexDevice SettingsPlexDevicesServersGet(ctx, )
Gets the user's available Plex servers

Returns a list of available Plex servers and their connectivity state

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]PlexDevice**](PlexDevice.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPlexGet**
> PlexSettings SettingsPlexGet(ctx, )
Get Plex settings

Retrieves current Plex settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PlexSettings**](PlexSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPlexLibraryGet**
> []PlexLibrary SettingsPlexLibraryGet(ctx, optional)
Get Plex libraries

Returns a list of Plex libraries in a JSON array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiSettingsPlexLibraryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiSettingsPlexLibraryGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sync** | **optional.String**| Syncs the current libraries with the current Plex server | 
 **enable** | **optional.String**| Comma separated list of libraries to enable. Any libraries not passed will be disabled! | 

### Return type

[**[]PlexLibrary**](PlexLibrary.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPlexPost**
> PlexSettings SettingsPlexPost(ctx, body)
Update Plex settings

Updates Plex settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PlexSettings**](PlexSettings.md)|  | 

### Return type

[**PlexSettings**](PlexSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPlexSyncGet**
> InlineResponse2002 SettingsPlexSyncGet(ctx, )
Get status of full Plex library scan

Returns scan progress in a JSON array.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPlexSyncPost**
> InlineResponse2002 SettingsPlexSyncPost(ctx, optional)
Start full Plex library scan

Runs a full Plex library scan and returns the progress in a JSON array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SettingsApiSettingsPlexSyncPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SettingsApiSettingsPlexSyncPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PlexSyncBody**](PlexSyncBody.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPlexUsersGet**
> []InlineResponse2003 SettingsPlexUsersGet(ctx, )
Get Plex users

Returns a list of Plex users in a JSON array.  Requires the `MANAGE_USERS` permission. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2003**](inline_response_200_3.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPublicGet**
> PublicSettings SettingsPublicGet(ctx, )
Get public settings

Returns settings that are not protected or sensitive. Mainly used to determine if the application has been configured for the first time.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PublicSettings**](PublicSettings.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsRadarrGet**
> []RadarrSettings SettingsRadarrGet(ctx, )
Get Radarr settings

Returns all Radarr settings in a JSON array.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsRadarrPost**
> RadarrSettings SettingsRadarrPost(ctx, body)
Create Radarr instance

Creates a new Radarr instance from the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RadarrSettings**](RadarrSettings.md)|  | 

### Return type

[**RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsRadarrRadarrIdDelete**
> RadarrSettings SettingsRadarrRadarrIdDelete(ctx, radarrId)
Delete Radarr instance

Deletes an existing Radarr instance based on the radarrId parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **radarrId** | **int32**| Radarr instance ID | 

### Return type

[**RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsRadarrRadarrIdProfilesGet**
> []ServiceProfile SettingsRadarrRadarrIdProfilesGet(ctx, radarrId)
Get available Radarr profiles

Returns a list of profiles available on the Radarr server instance in a JSON array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **radarrId** | **int32**| Radarr instance ID | 

### Return type

[**[]ServiceProfile**](ServiceProfile.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsRadarrRadarrIdPut**
> RadarrSettings SettingsRadarrRadarrIdPut(ctx, body, radarrId)
Update Radarr instance

Updates an existing Radarr instance with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RadarrSettings**](RadarrSettings.md)|  | 
  **radarrId** | **int32**| Radarr instance ID | 

### Return type

[**RadarrSettings**](RadarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsRadarrTestPost**
> InlineResponse2004 SettingsRadarrTestPost(ctx, body)
Test Radarr configuration

Tests if the Radarr configuration is valid. Returns profiles and root folders on success.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RadarrTestBody**](RadarrTestBody.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsSonarrGet**
> []SonarrSettings SettingsSonarrGet(ctx, )
Get Sonarr settings

Returns all Sonarr settings in a JSON array.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsSonarrPost**
> SonarrSettings SettingsSonarrPost(ctx, body)
Create Sonarr instance

Creates a new Sonarr instance from the request body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SonarrSettings**](SonarrSettings.md)|  | 

### Return type

[**SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsSonarrSonarrIdDelete**
> SonarrSettings SettingsSonarrSonarrIdDelete(ctx, sonarrId)
Delete Sonarr instance

Deletes an existing Sonarr instance based on the sonarrId parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sonarrId** | **int32**| Sonarr instance ID | 

### Return type

[**SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsSonarrSonarrIdPut**
> SonarrSettings SettingsSonarrSonarrIdPut(ctx, body, sonarrId)
Update Sonarr instance

Updates an existing Sonarr instance with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SonarrSettings**](SonarrSettings.md)|  | 
  **sonarrId** | **int32**| Sonarr instance ID | 

### Return type

[**SonarrSettings**](SonarrSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsSonarrTestPost**
> InlineResponse2004 SettingsSonarrTestPost(ctx, body)
Test Sonarr configuration

Tests if the Sonarr configuration is valid. Returns profiles and root folders on success.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SonarrTestBody**](SonarrTestBody.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsTautulliGet**
> TautulliSettings SettingsTautulliGet(ctx, )
Get Tautulli settings

Retrieves current Tautulli settings.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TautulliSettings**](TautulliSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsTautulliPost**
> TautulliSettings SettingsTautulliPost(ctx, body)
Update Tautulli settings

Updates Tautulli settings with the provided values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TautulliSettings**](TautulliSettings.md)|  | 

### Return type

[**TautulliSettings**](TautulliSettings.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

