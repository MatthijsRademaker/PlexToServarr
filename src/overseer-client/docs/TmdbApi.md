# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackdropsGet**](TmdbApi.md#BackdropsGet) | **Get** /backdrops | Get backdrops of trending items
[**GenresMovieGet**](TmdbApi.md#GenresMovieGet) | **Get** /genres/movie | Get list of official TMDB movie genres
[**GenresTvGet**](TmdbApi.md#GenresTvGet) | **Get** /genres/tv | Get list of official TMDB movie genres
[**LanguagesGet**](TmdbApi.md#LanguagesGet) | **Get** /languages | Languages supported by TMDB
[**NetworkNetworkIdGet**](TmdbApi.md#NetworkNetworkIdGet) | **Get** /network/{networkId} | Get TV network details
[**RegionsGet**](TmdbApi.md#RegionsGet) | **Get** /regions | Regions supported by TMDB
[**StudioStudioIdGet**](TmdbApi.md#StudioStudioIdGet) | **Get** /studio/{studioId} | Get movie studio details

# **BackdropsGet**
> []string BackdropsGet(ctx, )
Get backdrops of trending items

Returns a list of backdrop image paths in a JSON array.

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenresMovieGet**
> []InlineResponse20041 GenresMovieGet(ctx, optional)
Get list of official TMDB movie genres

Returns a list of genres in a JSON array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TmdbApiGenresMovieGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TmdbApiGenresMovieGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **optional.String**|  | 

### Return type

[**[]InlineResponse20041**](inline_response_200_41.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenresTvGet**
> []InlineResponse20042 GenresTvGet(ctx, optional)
Get list of official TMDB movie genres

Returns a list of genres in a JSON array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TmdbApiGenresTvGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TmdbApiGenresTvGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **optional.String**|  | 

### Return type

[**[]InlineResponse20042**](inline_response_200_42.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LanguagesGet**
> []InlineResponse20040 LanguagesGet(ctx, )
Languages supported by TMDB

Returns a list of languages in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse20040**](inline_response_200_40.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkNetworkIdGet**
> ProductionCompany NetworkNetworkIdGet(ctx, networkId)
Get TV network details

Returns TV network details in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **float64**|  | 

### Return type

[**ProductionCompany**](ProductionCompany.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegionsGet**
> []InlineResponse20039 RegionsGet(ctx, )
Regions supported by TMDB

Returns a list of regions in a JSON object.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse20039**](inline_response_200_39.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StudioStudioIdGet**
> ProductionCompany StudioStudioIdGet(ctx, studioId)
Get movie studio details

Returns movie studio details in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **studioId** | **float64**|  | 

### Return type

[**ProductionCompany**](ProductionCompany.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

