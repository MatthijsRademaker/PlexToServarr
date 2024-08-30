# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IssueCommentCommentIdDelete**](IssueApi.md#IssueCommentCommentIdDelete) | **Delete** /issueComment/{commentId} | Delete issue comment
[**IssueCommentCommentIdGet**](IssueApi.md#IssueCommentCommentIdGet) | **Get** /issueComment/{commentId} | Get issue comment
[**IssueCommentCommentIdPut**](IssueApi.md#IssueCommentCommentIdPut) | **Put** /issueComment/{commentId} | Update issue comment
[**IssueCountGet**](IssueApi.md#IssueCountGet) | **Get** /issue/count | Gets issue counts
[**IssueGet**](IssueApi.md#IssueGet) | **Get** /issue | Get all issues
[**IssueIssueIdCommentPost**](IssueApi.md#IssueIssueIdCommentPost) | **Post** /issue/{issueId}/comment | Create a comment
[**IssueIssueIdDelete**](IssueApi.md#IssueIssueIdDelete) | **Delete** /issue/{issueId} | Delete issue
[**IssueIssueIdGet**](IssueApi.md#IssueIssueIdGet) | **Get** /issue/{issueId} | Get issue
[**IssueIssueIdStatusPost**](IssueApi.md#IssueIssueIdStatusPost) | **Post** /issue/{issueId}/{status} | Update an issue&#x27;s status
[**IssuePost**](IssueApi.md#IssuePost) | **Post** /issue | Create new issue

# **IssueCommentCommentIdDelete**
> IssueCommentCommentIdDelete(ctx, commentId)
Delete issue comment

Deletes an issue comment. Only users with `MANAGE_ISSUES` or the user who created the comment can perform this action. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **string**| Issue Comment ID | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueCommentCommentIdGet**
> IssueComment IssueCommentCommentIdGet(ctx, commentId)
Get issue comment

Returns a single issue comment in JSON format. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **commentId** | **string**|  | 

### Return type

[**IssueComment**](IssueComment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueCommentCommentIdPut**
> IssueComment IssueCommentCommentIdPut(ctx, body, commentId)
Update issue comment

Updates and returns a single issue comment in JSON format. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IssueCommentCommentIdBody**](IssueCommentCommentIdBody.md)|  | 
  **commentId** | **string**|  | 

### Return type

[**IssueComment**](IssueComment.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueCountGet**
> InlineResponse20044 IssueCountGet(ctx, )
Gets issue counts

Returns the number of open and closed issues, as well as the number of issues of each type. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20044**](inline_response_200_44.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueGet**
> InlineResponse20043 IssueGet(ctx, optional)
Get all issues

Returns a list of issues in JSON format. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IssueApiIssueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueApiIssueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **optional.Float64**|  | 
 **skip** | **optional.Float64**|  | 
 **sort** | **optional.String**|  | [default to added]
 **filter** | **optional.String**|  | [default to open]
 **requestedBy** | **optional.Float64**|  | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueIssueIdCommentPost**
> Issue IssueIssueIdCommentPost(ctx, body, issueId)
Create a comment

Creates a comment and returns associated issue in JSON format. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IssueIdCommentBody**](IssueIdCommentBody.md)|  | 
  **issueId** | **float64**|  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueIssueIdDelete**
> IssueIssueIdDelete(ctx, issueId)
Delete issue

Removes an issue. If the user has the `MANAGE_ISSUES` permission, any issue can be removed. Otherwise, only a users own issues can be removed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| Issue ID | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueIssueIdGet**
> Issue IssueIssueIdGet(ctx, issueId)
Get issue

Returns a single issue in JSON format. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **float64**|  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssueIssueIdStatusPost**
> Issue IssueIssueIdStatusPost(ctx, issueId, status)
Update an issue's status

Updates an issue's status to approved or declined. Also returns the issue in a JSON object.  Requires the `MANAGE_ISSUES` permission or `ADMIN`. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **issueId** | **string**| Issue ID | 
  **status** | **string**| New status | 

### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IssuePost**
> Issue IssuePost(ctx, body)
Create new issue

Creates a new issue 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**IssueBody**](IssueBody.md)|  | 

### Return type

[**Issue**](Issue.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

