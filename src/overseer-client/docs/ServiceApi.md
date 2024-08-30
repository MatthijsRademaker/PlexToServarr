# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceRadarrGet**](ServiceApi.md#ServiceRadarrGet) | **Get** /service/radarr | Get non-sensitive Radarr server list
[**ServiceRadarrRadarrIdGet**](ServiceApi.md#ServiceRadarrRadarrIdGet) | **Get** /service/radarr/{radarrId} | Get Radarr server quality profiles and root folders
[**ServiceSonarrGet**](ServiceApi.md#ServiceSonarrGet) | **Get** /service/sonarr | Get non-sensitive Sonarr server list
[**ServiceSonarrLookupTmdbIdGet**](ServiceApi.md#ServiceSonarrLookupTmdbIdGet) | **Get** /service/sonarr/lookup/{tmdbId} | Get series from Sonarr
[**ServiceSonarrSonarrIdGet**](ServiceApi.md#ServiceSonarrSonarrIdGet) | **Get** /service/sonarr/{sonarrId} | Get Sonarr server quality profiles and root folders

# **ServiceRadarrGet**
> []RadarrSettings ServiceRadarrGet(ctx, )
Get non-sensitive Radarr server list

Returns a list of Radarr server IDs and names in a JSON object.

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

# **ServiceRadarrRadarrIdGet**
> InlineResponse20037 ServiceRadarrRadarrIdGet(ctx, radarrId)
Get Radarr server quality profiles and root folders

Returns a Radarr server's quality profile and root folder details in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **radarrId** | **float64**|  | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceSonarrGet**
> []SonarrSettings ServiceSonarrGet(ctx, )
Get non-sensitive Sonarr server list

Returns a list of Sonarr server IDs and names in a JSON object.

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

# **ServiceSonarrLookupTmdbIdGet**
> []SonarrSeries ServiceSonarrLookupTmdbIdGet(ctx, tmdbId)
Get series from Sonarr

Returns a list of series returned by searching for the name in Sonarr.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tmdbId** | **float64**|  | 

### Return type

[**[]SonarrSeries**](SonarrSeries.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceSonarrSonarrIdGet**
> InlineResponse20038 ServiceSonarrSonarrIdGet(ctx, sonarrId)
Get Sonarr server quality profiles and root folders

Returns a Sonarr server's quality profile and root folder details in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sonarrId** | **float64**|  | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

