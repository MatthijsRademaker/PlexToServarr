# {{classname}}

All URIs are relative to *{server}/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthMeGet**](UsersApi.md#AuthMeGet) | **Get** /auth/me | Get logged-in user
[**AuthResetPasswordGuidPost**](UsersApi.md#AuthResetPasswordGuidPost) | **Post** /auth/reset-password/{guid} | Reset the password for a user
[**AuthResetPasswordPost**](UsersApi.md#AuthResetPasswordPost) | **Post** /auth/reset-password | Send a reset password email
[**SettingsPlexUsersGet**](UsersApi.md#SettingsPlexUsersGet) | **Get** /settings/plex/users | Get Plex users
[**UserGet**](UsersApi.md#UserGet) | **Get** /user | Get all users
[**UserImportFromPlexPost**](UsersApi.md#UserImportFromPlexPost) | **Post** /user/import-from-plex | Import all users from Plex
[**UserPost**](UsersApi.md#UserPost) | **Post** /user | Create new user
[**UserPut**](UsersApi.md#UserPut) | **Put** /user | Update batch of users
[**UserRegisterPushSubscriptionPost**](UsersApi.md#UserRegisterPushSubscriptionPost) | **Post** /user/registerPushSubscription | Register a web push /user/registerPushSubscription
[**UserUserIdDelete**](UsersApi.md#UserUserIdDelete) | **Delete** /user/{userId} | Delete user by ID
[**UserUserIdGet**](UsersApi.md#UserUserIdGet) | **Get** /user/{userId} | Get user by ID
[**UserUserIdPut**](UsersApi.md#UserUserIdPut) | **Put** /user/{userId} | Update a user by user ID
[**UserUserIdQuotaGet**](UsersApi.md#UserUserIdQuotaGet) | **Get** /user/{userId}/quota | Get quotas for a specific user
[**UserUserIdRequestsGet**](UsersApi.md#UserUserIdRequestsGet) | **Get** /user/{userId}/requests | Get requests for a specific user
[**UserUserIdSettingsMainGet**](UsersApi.md#UserUserIdSettingsMainGet) | **Get** /user/{userId}/settings/main | Get general settings for a user
[**UserUserIdSettingsMainPost**](UsersApi.md#UserUserIdSettingsMainPost) | **Post** /user/{userId}/settings/main | Update general settings for a user
[**UserUserIdSettingsNotificationsGet**](UsersApi.md#UserUserIdSettingsNotificationsGet) | **Get** /user/{userId}/settings/notifications | Get notification settings for a user
[**UserUserIdSettingsNotificationsPost**](UsersApi.md#UserUserIdSettingsNotificationsPost) | **Post** /user/{userId}/settings/notifications | Update notification settings for a user
[**UserUserIdSettingsPasswordGet**](UsersApi.md#UserUserIdSettingsPasswordGet) | **Get** /user/{userId}/settings/password | Get password page informatiom
[**UserUserIdSettingsPasswordPost**](UsersApi.md#UserUserIdSettingsPasswordPost) | **Post** /user/{userId}/settings/password | Update password for a user
[**UserUserIdSettingsPermissionsGet**](UsersApi.md#UserUserIdSettingsPermissionsGet) | **Get** /user/{userId}/settings/permissions | Get permission settings for a user
[**UserUserIdSettingsPermissionsPost**](UsersApi.md#UserUserIdSettingsPermissionsPost) | **Post** /user/{userId}/settings/permissions | Update permission settings for a user
[**UserUserIdWatchDataGet**](UsersApi.md#UserUserIdWatchDataGet) | **Get** /user/{userId}/watch_data | Get watch data
[**UserUserIdWatchlistGet**](UsersApi.md#UserUserIdWatchlistGet) | **Get** /user/{userId}/watchlist | Get the Plex watchlist for a specific user

# **AuthMeGet**
> User AuthMeGet(ctx, )
Get logged-in user

Returns the currently logged-in user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthResetPasswordGuidPost**
> InlineResponse2009 AuthResetPasswordGuidPost(ctx, body, guid)
Reset the password for a user

Resets the password for a user if the given guid is connected to a user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ResetpasswordGuidBody**](ResetpasswordGuidBody.md)|  | 
  **guid** | **string**|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthResetPasswordPost**
> InlineResponse2009 AuthResetPasswordPost(ctx, body)
Send a reset password email

Sends a reset password email to the email if the user exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthResetpasswordBody**](AuthResetpasswordBody.md)|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsPlexUsersGet**
> []InlineResponse2003 SettingsPlexUsersGet(ctx, )
Get Plex users

Returns a list of Plex users in a JSON array.  Requires the `MANAGE_USERS` permission. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]InlineResponse2003**](inline_response_200_3.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserGet**
> InlineResponse20010 UserGet(ctx, optional)
Get all users

Returns all users in a JSON object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUserGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUserGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **take** | **optional.Float64**|  | 
 **skip** | **optional.Float64**|  | 
 **sort** | **optional.String**|  | [default to created]

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserImportFromPlexPost**
> []User UserImportFromPlexPost(ctx, optional)
Import all users from Plex

Fetches and imports users from the Plex server. If a list of Plex IDs is provided in the request body, only the specified users will be imported. Otherwise, all users will be imported.  Requires the `MANAGE_USERS` permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersApiUserImportFromPlexPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUserImportFromPlexPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UserImportfromplexBody**](UserImportfromplexBody.md)|  | 

### Return type

[**[]User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPost**
> User UserPost(ctx, body)
Create new user

Creates a new user. Requires the `MANAGE_USERS` permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserBody1**](UserBody1.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPut**
> []User UserPut(ctx, body)
Update batch of users

Update users with given IDs with provided values in request `body.settings`. You cannot update users' Plex tokens through this request.  Requires the `MANAGE_USERS` permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserBody**](UserBody.md)|  | 

### Return type

[**[]User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserRegisterPushSubscriptionPost**
> UserRegisterPushSubscriptionPost(ctx, body)
Register a web push /user/registerPushSubscription

Registers a web push subscription for the logged-in user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserRegisterPushSubscriptionBody**](UserRegisterPushSubscriptionBody.md)|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdDelete**
> User UserUserIdDelete(ctx, userId)
Delete user by ID

Deletes the user with the provided userId. Requires the `MANAGE_USERS` permission.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdGet**
> User UserUserIdGet(ctx, userId)
Get user by ID

Retrieves user details in a JSON object. Requires the `MANAGE_USERS` permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdPut**
> User UserUserIdPut(ctx, body, userId)
Update a user by user ID

Update a user with the provided values. You cannot update a user's Plex token through this request.  Requires the `MANAGE_USERS` permission. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**User**](User.md)|  | 
  **userId** | **float64**|  | 

### Return type

[**User**](User.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdQuotaGet**
> InlineResponse20012 UserUserIdQuotaGet(ctx, userId)
Get quotas for a specific user

Returns quota details for a user in a JSON object. Requires `MANAGE_USERS` permission if viewing other users. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdRequestsGet**
> InlineResponse20011 UserUserIdRequestsGet(ctx, userId, optional)
Get requests for a specific user

Retrieves a user's requests in a JSON object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 
 **optional** | ***UsersApiUserUserIdRequestsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUserUserIdRequestsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **take** | **optional.Float64**|  | 
 **skip** | **optional.Float64**|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdSettingsMainGet**
> InlineResponse20014 UserUserIdSettingsMainGet(ctx, userId)
Get general settings for a user

Returns general settings for a specific user. Requires `MANAGE_USERS` permission if viewing other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdSettingsMainPost**
> InlineResponse20014 UserUserIdSettingsMainPost(ctx, body, userId)
Update general settings for a user

Updates and returns general settings for a specific user. Requires `MANAGE_USERS` permission if editing other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsMainBody**](SettingsMainBody.md)|  | 
  **userId** | **float64**|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdSettingsNotificationsGet**
> UserSettingsNotifications UserUserIdSettingsNotificationsGet(ctx, userId)
Get notification settings for a user

Returns notification settings for a specific user. Requires `MANAGE_USERS` permission if viewing other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**UserSettingsNotifications**](UserSettingsNotifications.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdSettingsNotificationsPost**
> UserSettingsNotifications UserUserIdSettingsNotificationsPost(ctx, body, userId)
Update notification settings for a user

Updates and returns notification settings for a specific user. Requires `MANAGE_USERS` permission if editing other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserSettingsNotifications**](UserSettingsNotifications.md)|  | 
  **userId** | **float64**|  | 

### Return type

[**UserSettingsNotifications**](UserSettingsNotifications.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdSettingsPasswordGet**
> InlineResponse20015 UserUserIdSettingsPasswordGet(ctx, userId)
Get password page informatiom

Returns important data for the password page to function correctly. Requires `MANAGE_USERS` permission if viewing other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdSettingsPasswordPost**
> UserUserIdSettingsPasswordPost(ctx, body, userId)
Update password for a user

Updates a user's password. Requires `MANAGE_USERS` permission if editing other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsPasswordBody**](SettingsPasswordBody.md)|  | 
  **userId** | **float64**|  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdSettingsPermissionsGet**
> InlineResponse20016 UserUserIdSettingsPermissionsGet(ctx, userId)
Get permission settings for a user

Returns permission settings for a specific user. Requires `MANAGE_USERS` permission if viewing other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdSettingsPermissionsPost**
> InlineResponse20016 UserUserIdSettingsPermissionsPost(ctx, body, userId)
Update permission settings for a user

Updates and returns permission settings for a specific user. Requires `MANAGE_USERS` permission if editing other users.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsPermissionsBody**](SettingsPermissionsBody.md)|  | 
  **userId** | **float64**|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdWatchDataGet**
> InlineResponse20017 UserUserIdWatchDataGet(ctx, userId)
Get watch data

Returns play count, play duration, and recently watched media.  Requires the `ADMIN` permission to fetch results for other users. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUserIdWatchlistGet**
> InlineResponse20013 UserUserIdWatchlistGet(ctx, userId, optional)
Get the Plex watchlist for a specific user

Retrieves a user's Plex Watchlist in a JSON object. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **float64**|  | 
 **optional** | ***UsersApiUserUserIdWatchlistGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersApiUserUserIdWatchlistGetOpts struct
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

