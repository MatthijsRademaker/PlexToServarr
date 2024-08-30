# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TvTvIdGet**](TvApi.md#TvTvIdGet) | **Get** /tv/{tvId} | Get TV details
[**TvTvIdRatingsGet**](TvApi.md#TvTvIdRatingsGet) | **Get** /tv/{tvId}/ratings | Get TV ratings
[**TvTvIdRecommendationsGet**](TvApi.md#TvTvIdRecommendationsGet) | **Get** /tv/{tvId}/recommendations | Get recommended TV series
[**TvTvIdSeasonSeasonIdGet**](TvApi.md#TvTvIdSeasonSeasonIdGet) | **Get** /tv/{tvId}/season/{seasonId} | Get season details and episode list
[**TvTvIdSimilarGet**](TvApi.md#TvTvIdSimilarGet) | **Get** /tv/{tvId}/similar | Get similar TV series

# **TvTvIdGet**
> TvDetails TvTvIdGet(ctx, tvId, optional)
Get TV details

Returns full TV details in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tvId** | **float64**|  | 
 **optional** | ***TvApiTvTvIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TvApiTvTvIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **optional.String**|  | 

### Return type

[**TvDetails**](TvDetails.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TvTvIdRatingsGet**
> InlineResponse20033 TvTvIdRatingsGet(ctx, tvId)
Get TV ratings

Returns ratings based on provided tvId in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tvId** | **float64**|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TvTvIdRecommendationsGet**
> InlineResponse20025 TvTvIdRecommendationsGet(ctx, tvId, optional)
Get recommended TV series

Returns list of recommended TV series based on the provided tvId in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tvId** | **float64**|  | 
 **optional** | ***TvApiTvTvIdRecommendationsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TvApiTvTvIdRecommendationsGetOpts struct
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

# **TvTvIdSeasonSeasonIdGet**
> Season TvTvIdSeasonSeasonIdGet(ctx, tvId, seasonId, optional)
Get season details and episode list

Returns season details with a list of episodes in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tvId** | **float64**|  | 
  **seasonId** | **float64**|  | 
 **optional** | ***TvApiTvTvIdSeasonSeasonIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TvApiTvTvIdSeasonSeasonIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **language** | **optional.String**|  | 

### Return type

[**Season**](Season.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TvTvIdSimilarGet**
> InlineResponse20025 TvTvIdSimilarGet(ctx, tvId, optional)
Get similar TV series

Returns list of similar TV series based on the provided tvId in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tvId** | **float64**|  | 
 **optional** | ***TvApiTvTvIdSimilarGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TvApiTvTvIdSimilarGetOpts struct
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

