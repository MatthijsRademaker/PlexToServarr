# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MediaGet**](MediaApi.md#MediaGet) | **Get** /media | Get media
[**MediaMediaIdDelete**](MediaApi.md#MediaMediaIdDelete) | **Delete** /media/{mediaId} | Delete media item
[**MediaMediaIdStatusPost**](MediaApi.md#MediaMediaIdStatusPost) | **Post** /media/{mediaId}/{status} | Update media status
[**MediaMediaIdWatchDataGet**](MediaApi.md#MediaMediaIdWatchDataGet) | **Get** /media/{mediaId}/watch_data | Get watch data

# **MediaGet**
> InlineResponse20035 MediaGet(ctx, optional)
Get media

Returns all media (can be filtered and limited) in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MediaApiMediaGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MediaApiMediaGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **optional.Float64**|  | 
 **skip** | **optional.Float64**|  | 
 **filter** | **optional.String**|  | 
 **sort** | **optional.String**|  | [default to added]

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaMediaIdDelete**
> MediaMediaIdDelete(ctx, mediaId)
Delete media item

Removes a media item. The `MANAGE_REQUESTS` permission is required to perform this action.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaId** | **string**| Media ID | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaMediaIdStatusPost**
> MediaInfo MediaMediaIdStatusPost(ctx, mediaId, status, optional)
Update media status

Updates a media item's status and returns the media in JSON format

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaId** | **string**| Media ID | 
  **status** | **string**| New status | 
 **optional** | ***MediaApiMediaMediaIdStatusPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MediaApiMediaMediaIdStatusPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MediaIdStatusBody**](MediaIdStatusBody.md)|  | 

### Return type

[**MediaInfo**](MediaInfo.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MediaMediaIdWatchDataGet**
> InlineResponse20036 MediaMediaIdWatchDataGet(ctx, mediaId)
Get watch data

Returns play count, play duration, and users who have watched the media.  Requires the `ADMIN` permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mediaId** | **string**| Media ID | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

