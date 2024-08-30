# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverGenresliderMovieGet**](SearchApi.md#DiscoverGenresliderMovieGet) | **Get** /discover/genreslider/movie | Get genre slider data for movies
[**DiscoverGenresliderTvGet**](SearchApi.md#DiscoverGenresliderTvGet) | **Get** /discover/genreslider/tv | Get genre slider data for TV series
[**DiscoverKeywordKeywordIdMoviesGet**](SearchApi.md#DiscoverKeywordKeywordIdMoviesGet) | **Get** /discover/keyword/{keywordId}/movies | Get movies from keyword
[**DiscoverMoviesGenreGenreIdGet**](SearchApi.md#DiscoverMoviesGenreGenreIdGet) | **Get** /discover/movies/genre/{genreId} | Discover movies by genre
[**DiscoverMoviesGet**](SearchApi.md#DiscoverMoviesGet) | **Get** /discover/movies | Discover movies
[**DiscoverMoviesLanguageLanguageGet**](SearchApi.md#DiscoverMoviesLanguageLanguageGet) | **Get** /discover/movies/language/{language} | Discover movies by original language
[**DiscoverMoviesStudioStudioIdGet**](SearchApi.md#DiscoverMoviesStudioStudioIdGet) | **Get** /discover/movies/studio/{studioId} | Discover movies by studio
[**DiscoverMoviesUpcomingGet**](SearchApi.md#DiscoverMoviesUpcomingGet) | **Get** /discover/movies/upcoming | Upcoming movies
[**DiscoverTrendingGet**](SearchApi.md#DiscoverTrendingGet) | **Get** /discover/trending | Trending movies and TV
[**DiscoverTvGenreGenreIdGet**](SearchApi.md#DiscoverTvGenreGenreIdGet) | **Get** /discover/tv/genre/{genreId} | Discover TV shows by genre
[**DiscoverTvGet**](SearchApi.md#DiscoverTvGet) | **Get** /discover/tv | Discover TV shows
[**DiscoverTvLanguageLanguageGet**](SearchApi.md#DiscoverTvLanguageLanguageGet) | **Get** /discover/tv/language/{language} | Discover TV shows by original language
[**DiscoverTvNetworkNetworkIdGet**](SearchApi.md#DiscoverTvNetworkNetworkIdGet) | **Get** /discover/tv/network/{networkId} | Discover TV shows by network
[**DiscoverTvUpcomingGet**](SearchApi.md#DiscoverTvUpcomingGet) | **Get** /discover/tv/upcoming | Discover Upcoming TV shows
[**DiscoverWatchlistGet**](SearchApi.md#DiscoverWatchlistGet) | **Get** /discover/watchlist | Get the Plex watchlist.
[**SearchCompanyGet**](SearchApi.md#SearchCompanyGet) | **Get** /search/company | Search for companies
[**SearchGet**](SearchApi.md#SearchGet) | **Get** /search | Search for movies, TV shows, or people
[**SearchKeywordGet**](SearchApi.md#SearchKeywordGet) | **Get** /search/keyword | Search for keywords

# **DiscoverGenresliderMovieGet**
> []InlineResponse20029 DiscoverGenresliderMovieGet(ctx, optional)
Get genre slider data for movies

Returns a list of genres with backdrops attached

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiDiscoverGenresliderMovieGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverGenresliderMovieGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **optional.String**|  | 

### Return type

[**[]InlineResponse20029**](inline_response_200_29.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverGenresliderTvGet**
> []InlineResponse20029 DiscoverGenresliderTvGet(ctx, optional)
Get genre slider data for TV series

Returns a list of genres with backdrops attached

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiDiscoverGenresliderTvGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverGenresliderTvGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **language** | **optional.String**|  | 

### Return type

[**[]InlineResponse20029**](inline_response_200_29.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverKeywordKeywordIdMoviesGet**
> InlineResponse20021 DiscoverKeywordKeywordIdMoviesGet(ctx, keywordId, optional)
Get movies from keyword

Returns list of movies based on the provided keyword ID a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keywordId** | **float64**|  | 
 **optional** | ***SearchApiDiscoverKeywordKeywordIdMoviesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverKeywordKeywordIdMoviesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverMoviesGenreGenreIdGet**
> InlineResponse20022 DiscoverMoviesGenreGenreIdGet(ctx, genreId, optional)
Discover movies by genre

Returns a list of movies based on the provided genre ID in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **genreId** | **string**|  | 
 **optional** | ***SearchApiDiscoverMoviesGenreGenreIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverMoviesGenreGenreIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverMoviesGet**
> InlineResponse20021 DiscoverMoviesGet(ctx, optional)
Discover movies

Returns a list of movies in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiDiscoverMoviesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverMoviesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 
 **genre** | **optional.String**|  | 
 **studio** | **optional.Float64**|  | 
 **keywords** | **optional.String**|  | 
 **sortBy** | **optional.String**|  | 
 **primaryReleaseDateGte** | **optional.String**|  | 
 **primaryReleaseDateLte** | **optional.String**|  | 
 **withRuntimeGte** | **optional.Float64**|  | 
 **withRuntimeLte** | **optional.Float64**|  | 
 **voteAverageGte** | **optional.Float64**|  | 
 **voteAverageLte** | **optional.Float64**|  | 
 **voteCountGte** | **optional.Float64**|  | 
 **voteCountLte** | **optional.Float64**|  | 
 **watchRegion** | **optional.String**|  | 
 **watchProviders** | **optional.String**|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverMoviesLanguageLanguageGet**
> InlineResponse20023 DiscoverMoviesLanguageLanguageGet(ctx, language, optional)
Discover movies by original language

Returns a list of movies based on the provided ISO 639-1 language code in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **language** | **string**|  | 
 **optional** | ***SearchApiDiscoverMoviesLanguageLanguageGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverMoviesLanguageLanguageGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverMoviesStudioStudioIdGet**
> InlineResponse20024 DiscoverMoviesStudioStudioIdGet(ctx, studioId, optional)
Discover movies by studio

Returns a list of movies based on the provided studio ID in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **studioId** | **string**|  | 
 **optional** | ***SearchApiDiscoverMoviesStudioStudioIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverMoviesStudioStudioIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverMoviesUpcomingGet**
> InlineResponse20021 DiscoverMoviesUpcomingGet(ctx, optional)
Upcoming movies

Returns a list of movies in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiDiscoverMoviesUpcomingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverMoviesUpcomingGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverTrendingGet**
> InlineResponse20018 DiscoverTrendingGet(ctx, optional)
Trending movies and TV

Returns a list of movies and TV shows in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiDiscoverTrendingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverTrendingGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverTvGenreGenreIdGet**
> InlineResponse20027 DiscoverTvGenreGenreIdGet(ctx, genreId, optional)
Discover TV shows by genre

Returns a list of TV shows based on the provided genre ID in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **genreId** | **string**|  | 
 **optional** | ***SearchApiDiscoverTvGenreGenreIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverTvGenreGenreIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverTvGet**
> InlineResponse20025 DiscoverTvGet(ctx, optional)
Discover TV shows

Returns a list of TV shows in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiDiscoverTvGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverTvGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 
 **genre** | **optional.String**|  | 
 **network** | **optional.Float64**|  | 
 **keywords** | **optional.String**|  | 
 **sortBy** | **optional.String**|  | 
 **firstAirDateGte** | **optional.String**|  | 
 **firstAirDateLte** | **optional.String**|  | 
 **withRuntimeGte** | **optional.Float64**|  | 
 **withRuntimeLte** | **optional.Float64**|  | 
 **voteAverageGte** | **optional.Float64**|  | 
 **voteAverageLte** | **optional.Float64**|  | 
 **voteCountGte** | **optional.Float64**|  | 
 **voteCountLte** | **optional.Float64**|  | 
 **watchRegion** | **optional.String**|  | 
 **watchProviders** | **optional.String**|  | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverTvLanguageLanguageGet**
> InlineResponse20026 DiscoverTvLanguageLanguageGet(ctx, language, optional)
Discover TV shows by original language

Returns a list of TV shows based on the provided ISO 639-1 language code in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **language** | **string**|  | 
 **optional** | ***SearchApiDiscoverTvLanguageLanguageGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverTvLanguageLanguageGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverTvNetworkNetworkIdGet**
> InlineResponse20028 DiscoverTvNetworkNetworkIdGet(ctx, networkId, optional)
Discover TV shows by network

Returns a list of TV shows based on the provided network ID in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **networkId** | **string**|  | 
 **optional** | ***SearchApiDiscoverTvNetworkNetworkIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverTvNetworkNetworkIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverTvUpcomingGet**
> InlineResponse20025 DiscoverTvUpcomingGet(ctx, optional)
Discover Upcoming TV shows

Returns a list of upcoming TV shows in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiDiscoverTvUpcomingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverTvUpcomingGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DiscoverWatchlistGet**
> InlineResponse20013 DiscoverWatchlistGet(ctx, optional)
Get the Plex watchlist.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchApiDiscoverWatchlistGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiDiscoverWatchlistGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Float64**|  | [default to 1]

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCompanyGet**
> InlineResponse20020 SearchCompanyGet(ctx, query, optional)
Search for companies

Returns a list of TMDB companies matching the search query. (Will not return origin country)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**|  | 
 **optional** | ***SearchApiSearchCompanyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchCompanyGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchGet**
> InlineResponse20018 SearchGet(ctx, query, optional)
Search for movies, TV shows, or people

Returns a list of movies, TV shows, or people a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**|  | 
 **optional** | ***SearchApiSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]
 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchKeywordGet**
> InlineResponse20019 SearchKeywordGet(ctx, query, optional)
Search for keywords

Returns a list of TMDB keywords matching the search query

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**|  | 
 **optional** | ***SearchApiSearchKeywordGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchApiSearchKeywordGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float64**|  | [default to 1]

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

