# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeywordKeywordIdGet**](OtherApi.md#KeywordKeywordIdGet) | **Get** /keyword/{keywordId} | Get keyword
[**WatchprovidersMoviesGet**](OtherApi.md#WatchprovidersMoviesGet) | **Get** /watchproviders/movies | Get watch provider movies
[**WatchprovidersRegionsGet**](OtherApi.md#WatchprovidersRegionsGet) | **Get** /watchproviders/regions | Get watch provider regions
[**WatchprovidersTvGet**](OtherApi.md#WatchprovidersTvGet) | **Get** /watchproviders/tv | Get watch provider series

# **KeywordKeywordIdGet**
> Keyword KeywordKeywordIdGet(ctx, keywordId)
Get keyword

Returns a single keyword in JSON format. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **keywordId** | **float64**|  | 

### Return type

[**Keyword**](Keyword.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchprovidersMoviesGet**
> []WatchProviderDetails WatchprovidersMoviesGet(ctx, watchRegion)
Get watch provider movies

Returns a list of all available watch providers for movies. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **watchRegion** | **string**|  | 

### Return type

[**[]WatchProviderDetails**](WatchProviderDetails.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchprovidersRegionsGet**
> []WatchProviderRegion WatchprovidersRegionsGet(ctx, )
Get watch provider regions

Returns a list of all available watch provider regions. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]WatchProviderRegion**](WatchProviderRegion.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchprovidersTvGet**
> []WatchProviderDetails WatchprovidersTvGet(ctx, watchRegion)
Get watch provider series

Returns a list of all available watch providers for series. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **watchRegion** | **string**|  | 

### Return type

[**[]WatchProviderDetails**](WatchProviderDetails.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

