/*
 * Radarr
 *
 * Radarr API docs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientRadarr

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type HistoryApiService service

/*
HistoryApiService
  - @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
  - @param id
*/
func (a *HistoryApiService) ApiV3HistoryFailedIdPost(ctx context.Context, id int32) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/history/failed/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Api-Key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
HistoryApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *HistoryApiApiV3HistoryGetOpts - Optional Parameters:
     * @param "Page" (optional.Int32) -
     * @param "PageSize" (optional.Int32) -
     * @param "SortKey" (optional.String) -
     * @param "SortDirection" (optional.Interface of SortDirection) -
     * @param "IncludeMovie" (optional.Bool) -
     * @param "EventType" (optional.Interface of []int32) -
     * @param "DownloadId" (optional.String) -
     * @param "MovieIds" (optional.Interface of []int32) -
     * @param "Languages" (optional.Interface of []int32) -
     * @param "Quality" (optional.Interface of []int32) -
@return HistoryResourcePagingResource
*/

type HistoryApiApiV3HistoryGetOpts struct {
	Page          optional.Int32
	PageSize      optional.Int32
	SortKey       optional.String
	SortDirection optional.Interface
	IncludeMovie  optional.Bool
	EventType     optional.Interface
	DownloadId    optional.String
	MovieIds      optional.Interface
	Languages     optional.Interface
	Quality       optional.Interface
}

func (a *HistoryApiService) ApiV3HistoryGet(ctx context.Context, localVarOptionals *HistoryApiApiV3HistoryGetOpts) (HistoryResourcePagingResource, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue HistoryResourcePagingResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/history"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PageSize.IsSet() {
		localVarQueryParams.Add("pageSize", parameterToString(localVarOptionals.PageSize.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortKey.IsSet() {
		localVarQueryParams.Add("sortKey", parameterToString(localVarOptionals.SortKey.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SortDirection.IsSet() {
		localVarQueryParams.Add("sortDirection", parameterToString(localVarOptionals.SortDirection.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeMovie.IsSet() {
		localVarQueryParams.Add("includeMovie", parameterToString(localVarOptionals.IncludeMovie.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EventType.IsSet() {
		localVarQueryParams.Add("eventType", parameterToString(localVarOptionals.EventType.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.DownloadId.IsSet() {
		localVarQueryParams.Add("downloadId", parameterToString(localVarOptionals.DownloadId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MovieIds.IsSet() {
		localVarQueryParams.Add("movieIds", parameterToString(localVarOptionals.MovieIds.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Languages.IsSet() {
		localVarQueryParams.Add("languages", parameterToString(localVarOptionals.Languages.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Quality.IsSet() {
		localVarQueryParams.Add("quality", parameterToString(localVarOptionals.Quality.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Api-Key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v HistoryResourcePagingResource
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
HistoryApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *HistoryApiApiV3HistoryMovieGetOpts - Optional Parameters:
     * @param "MovieId" (optional.Int32) -
     * @param "EventType" (optional.Interface of MovieHistoryEventType) -
     * @param "IncludeMovie" (optional.Bool) -
@return []HistoryResource
*/

type HistoryApiApiV3HistoryMovieGetOpts struct {
	MovieId      optional.Int32
	EventType    optional.Interface
	IncludeMovie optional.Bool
}

func (a *HistoryApiService) ApiV3HistoryMovieGet(ctx context.Context, localVarOptionals *HistoryApiApiV3HistoryMovieGetOpts) ([]HistoryResource, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []HistoryResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/history/movie"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.MovieId.IsSet() {
		localVarQueryParams.Add("movieId", parameterToString(localVarOptionals.MovieId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EventType.IsSet() {
		localVarQueryParams.Add("eventType", parameterToString(localVarOptionals.EventType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeMovie.IsSet() {
		localVarQueryParams.Add("includeMovie", parameterToString(localVarOptionals.IncludeMovie.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Api-Key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []HistoryResource
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
HistoryApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *HistoryApiApiV3HistorySinceGetOpts - Optional Parameters:
     * @param "Date" (optional.Time) -
     * @param "EventType" (optional.Interface of MovieHistoryEventType) -
     * @param "IncludeMovie" (optional.Bool) -
@return []HistoryResource
*/

type HistoryApiApiV3HistorySinceGetOpts struct {
	Date         optional.Time
	EventType    optional.Interface
	IncludeMovie optional.Bool
}

func (a *HistoryApiService) ApiV3HistorySinceGet(ctx context.Context, localVarOptionals *HistoryApiApiV3HistorySinceGetOpts) ([]HistoryResource, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []HistoryResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/history/since"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Date.IsSet() {
		localVarQueryParams.Add("date", parameterToString(localVarOptionals.Date.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EventType.IsSet() {
		localVarQueryParams.Add("eventType", parameterToString(localVarOptionals.EventType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeMovie.IsSet() {
		localVarQueryParams.Add("includeMovie", parameterToString(localVarOptionals.IncludeMovie.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["X-Api-Key"] = key

		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []HistoryResource
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}