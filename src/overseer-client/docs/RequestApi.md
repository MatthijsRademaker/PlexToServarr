# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestCountGet**](RequestApi.md#RequestCountGet) | **Get** /request/count | Gets request counts
[**RequestGet**](RequestApi.md#RequestGet) | **Get** /request | Get all requests
[**RequestPost**](RequestApi.md#RequestPost) | **Post** /request | Create new request
[**RequestRequestIdDelete**](RequestApi.md#RequestRequestIdDelete) | **Delete** /request/{requestId} | Delete request
[**RequestRequestIdGet**](RequestApi.md#RequestRequestIdGet) | **Get** /request/{requestId} | Get MediaRequest
[**RequestRequestIdPut**](RequestApi.md#RequestRequestIdPut) | **Put** /request/{requestId} | Update MediaRequest
[**RequestRequestIdRetryPost**](RequestApi.md#RequestRequestIdRetryPost) | **Post** /request/{requestId}/retry | Retry failed request
[**RequestRequestIdStatusPost**](RequestApi.md#RequestRequestIdStatusPost) | **Post** /request/{requestId}/{status} | Update a request&#x27;s status

# **RequestCountGet**
> InlineResponse20030 RequestCountGet(ctx, )
Gets request counts

Returns the number of pending and approved requests. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestGet**
> InlineResponse20011 RequestGet(ctx, optional)
Get all requests

Returns all requests if the user has the `ADMIN` or `MANAGE_REQUESTS` permissions. Otherwise, only the logged-in user's requests are returned.  If the `requestedBy` parameter is specified, only requests from that particular user ID will be returned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RequestApiRequestGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequestApiRequestGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **optional.Float64**|  | 
 **skip** | **optional.Float64**|  | 
 **filter** | **optional.String**|  | 
 **sort** | **optional.String**|  | [default to added]
 **requestedBy** | **optional.Float64**|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestPost**
> MediaRequest RequestPost(ctx, body)
Create new request

Creates a new request with the provided media ID and type. The `REQUEST` permission is required.  If the user has the `ADMIN` or `AUTO_APPROVE` permissions, their request will be auomatically approved. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestBody**](RequestBody.md)|  | 

### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRequestIdDelete**
> RequestRequestIdDelete(ctx, requestId)
Delete request

Removes a request. If the user has the `MANAGE_REQUESTS` permission, any request can be removed. Otherwise, only pending requests can be removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Request ID | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRequestIdGet**
> MediaRequest RequestRequestIdGet(ctx, requestId)
Get MediaRequest

Returns a specific MediaRequest in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Request ID | 

### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRequestIdPut**
> MediaRequest RequestRequestIdPut(ctx, body, requestId)
Update MediaRequest

Updates a specific media request and returns the request in a JSON object. Requires the `MANAGE_REQUESTS` permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RequestRequestIdBody**](RequestRequestIdBody.md)|  | 
  **requestId** | **string**| Request ID | 

### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRequestIdRetryPost**
> MediaRequest RequestRequestIdRetryPost(ctx, requestId)
Retry failed request

Retries a request by resending requests to Sonarr or Radarr.  Requires the `MANAGE_REQUESTS` permission or `ADMIN`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Request ID | 

### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestRequestIdStatusPost**
> MediaRequest RequestRequestIdStatusPost(ctx, requestId, status)
Update a request's status

Updates a request's status to approved or declined. Also returns the request in a JSON object.  Requires the `MANAGE_REQUESTS` permission or `ADMIN`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| Request ID | 
  **status** | **string**| New status | 

### Return type

[**MediaRequest**](MediaRequest.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

