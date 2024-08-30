# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MovieMovieIdGet**](MoviesApi.md#MovieMovieIdGet) | **Get** /movie/{movieId} | Get movie details
[**MovieMovieIdRatingsGet**](MoviesApi.md#MovieMovieIdRatingsGet) | **Get** /movie/{movieId}/ratings | Get movie ratings
[**MovieMovieIdRatingscombinedGet**](MoviesApi.md#MovieMovieIdRatingscombinedGet) | **Get** /movie/{movieId}/ratingscombined | Get RT and IMDB movie ratings combined
[**MovieMovieIdRecommendationsGet**](MoviesApi.md#MovieMovieIdRecommendationsGet) | **Get** /movie/{movieId}/recommendations | Get recommended movies
[**MovieMovieIdSimilarGet**](MoviesApi.md#MovieMovieIdSimilarGet) | **Get** /movie/{movieId}/similar | Get similar movies

# **MovieMovieIdGet**
> MovieDetails MovieMovieIdGet(ctx, movieId, optional)
Get movie details

Returns full movie details in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **movieId** | **float64**|  | 
 **optional** | ***MoviesApiMovieMovieIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MoviesApiMovieMovieIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **optional.String**|  | 

### Return type

[**MovieDetails**](MovieDetails.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MovieMovieIdRatingsGet**
> InlineResponse20031 MovieMovieIdRatingsGet(ctx, movieId)
Get movie ratings

Returns ratings based on the provided movieId in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **movieId** | **float64**|  | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MovieMovieIdRatingscombinedGet**
> InlineResponse20032 MovieMovieIdRatingscombinedGet(ctx, movieId)
Get RT and IMDB movie ratings combined

Returns ratings from RottenTomatoes and IMDB based on the provided movieId in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **movieId** | **float64**|  | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MovieMovieIdRecommendationsGet**
> InlineResponse20021 MovieMovieIdRecommendationsGet(ctx, movieId, optional)
Get recommended movies

Returns list of recommended movies based on provided movie ID in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **movieId** | **float64**|  | 
 **optional** | ***MoviesApiMovieMovieIdRecommendationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MoviesApiMovieMovieIdRecommendationsGetOpts struct
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

# **MovieMovieIdSimilarGet**
> InlineResponse20021 MovieMovieIdSimilarGet(ctx, movieId, optional)
Get similar movies

Returns list of similar movies based on the provided movieId in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **movieId** | **float64**|  | 
 **optional** | ***MoviesApiMovieMovieIdSimilarGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MoviesApiMovieMovieIdSimilarGetOpts struct
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

