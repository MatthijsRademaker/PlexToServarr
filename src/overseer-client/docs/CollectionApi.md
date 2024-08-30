# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CollectionCollectionIdGet**](CollectionApi.md#CollectionCollectionIdGet) | **Get** /collection/{collectionId} | Get collection details

# **CollectionCollectionIdGet**
> Collection CollectionCollectionIdGet(ctx, collectionId, optional)
Get collection details

Returns full collection details in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **collectionId** | **float64**|  | 
 **optional** | ***CollectionApiCollectionCollectionIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CollectionApiCollectionCollectionIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **optional.String**|  | 

### Return type

[**Collection**](Collection.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

