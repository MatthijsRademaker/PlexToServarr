# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PersonPersonIdCombinedCreditsGet**](PersonApi.md#PersonPersonIdCombinedCreditsGet) | **Get** /person/{personId}/combined_credits | Get combined credits
[**PersonPersonIdGet**](PersonApi.md#PersonPersonIdGet) | **Get** /person/{personId} | Get person details

# **PersonPersonIdCombinedCreditsGet**
> InlineResponse20034 PersonPersonIdCombinedCreditsGet(ctx, personId, optional)
Get combined credits

Returns the person's combined credits based on the provided personId in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **personId** | **float64**|  | 
 **optional** | ***PersonApiPersonPersonIdCombinedCreditsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonApiPersonPersonIdCombinedCreditsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **optional.String**|  | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PersonPersonIdGet**
> PersonDetails PersonPersonIdGet(ctx, personId, optional)
Get person details

Returns person details based on provided personId in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **personId** | **float64**|  | 
 **optional** | ***PersonApiPersonPersonIdGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonApiPersonPersonIdGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **language** | **optional.String**|  | 

### Return type

[**PersonDetails**](PersonDetails.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

