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

type QueueApiService service

/*
QueueApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *QueueApiApiV3QueueBulkDeleteOpts - Optional Parameters:
     * @param "Body" (optional.Interface of QueueBulkResource) -
     * @param "RemoveFromClient" (optional.Bool) -
     * @param "Blocklist" (optional.Bool) -
     * @param "SkipRedownload" (optional.Bool) -
     * @param "ChangeCategory" (optional.Bool) -

*/

type QueueApiApiV3QueueBulkDeleteOpts struct {
	Body             optional.Interface
	RemoveFromClient optional.Bool
	Blocklist        optional.Bool
	SkipRedownload   optional.Bool
	ChangeCategory   optional.Bool
}

func (a *QueueApiService) ApiV3QueueBulkDelete(ctx context.Context, localVarOptionals *QueueApiApiV3QueueBulkDeleteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/queue/bulk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.RemoveFromClient.IsSet() {
		localVarQueryParams.Add("removeFromClient", parameterToString(localVarOptionals.RemoveFromClient.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Blocklist.IsSet() {
		localVarQueryParams.Add("blocklist", parameterToString(localVarOptionals.Blocklist.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SkipRedownload.IsSet() {
		localVarQueryParams.Add("skipRedownload", parameterToString(localVarOptionals.SkipRedownload.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChangeCategory.IsSet() {
		localVarQueryParams.Add("changeCategory", parameterToString(localVarOptionals.ChangeCategory.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "text/json", "application/_*+json"}

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
	// body params
	if localVarOptionals != nil && localVarOptionals.Body.IsSet() {

		localVarOptionalBody := localVarOptionals.Body.Value()
		localVarPostBody = &localVarOptionalBody
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
QueueApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *QueueApiApiV3QueueGetOpts - Optional Parameters:
     * @param "Page" (optional.Int32) -
     * @param "PageSize" (optional.Int32) -
     * @param "SortKey" (optional.String) -
     * @param "SortDirection" (optional.Interface of SortDirection) -
     * @param "IncludeUnknownMovieItems" (optional.Bool) -
     * @param "IncludeMovie" (optional.Bool) -
     * @param "MovieIds" (optional.Interface of []int32) -
     * @param "Protocol" (optional.Interface of DownloadProtocol) -
     * @param "Languages" (optional.Interface of []int32) -
     * @param "Quality" (optional.Int32) -
@return QueueResourcePagingResource
*/

type QueueApiApiV3QueueGetOpts struct {
	Page                     optional.Int32
	PageSize                 optional.Int32
	SortKey                  optional.String
	SortDirection            optional.Interface
	IncludeUnknownMovieItems optional.Bool
	IncludeMovie             optional.Bool
	MovieIds                 optional.Interface
	Protocol                 optional.Interface
	Languages                optional.Interface
	Quality                  optional.Int32
}

func (a *QueueApiService) ApiV3QueueGet(ctx context.Context, localVarOptionals *QueueApiApiV3QueueGetOpts) (QueueResourcePagingResource, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue QueueResourcePagingResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/queue"

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
	if localVarOptionals != nil && localVarOptionals.IncludeUnknownMovieItems.IsSet() {
		localVarQueryParams.Add("includeUnknownMovieItems", parameterToString(localVarOptionals.IncludeUnknownMovieItems.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeMovie.IsSet() {
		localVarQueryParams.Add("includeMovie", parameterToString(localVarOptionals.IncludeMovie.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MovieIds.IsSet() {
		localVarQueryParams.Add("movieIds", parameterToString(localVarOptionals.MovieIds.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Protocol.IsSet() {
		localVarQueryParams.Add("protocol", parameterToString(localVarOptionals.Protocol.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Languages.IsSet() {
		localVarQueryParams.Add("languages", parameterToString(localVarOptionals.Languages.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Quality.IsSet() {
		localVarQueryParams.Add("quality", parameterToString(localVarOptionals.Quality.Value(), ""))
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
			var v QueueResourcePagingResource
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
QueueApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @param optional nil or *QueueApiApiV3QueueIdDeleteOpts - Optional Parameters:
     * @param "RemoveFromClient" (optional.Bool) -
     * @param "Blocklist" (optional.Bool) -
     * @param "SkipRedownload" (optional.Bool) -
     * @param "ChangeCategory" (optional.Bool) -

*/

type QueueApiApiV3QueueIdDeleteOpts struct {
	RemoveFromClient optional.Bool
	Blocklist        optional.Bool
	SkipRedownload   optional.Bool
	ChangeCategory   optional.Bool
}

func (a *QueueApiService) ApiV3QueueIdDelete(ctx context.Context, id int32, localVarOptionals *QueueApiApiV3QueueIdDeleteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/queue/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.RemoveFromClient.IsSet() {
		localVarQueryParams.Add("removeFromClient", parameterToString(localVarOptionals.RemoveFromClient.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Blocklist.IsSet() {
		localVarQueryParams.Add("blocklist", parameterToString(localVarOptionals.Blocklist.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SkipRedownload.IsSet() {
		localVarQueryParams.Add("skipRedownload", parameterToString(localVarOptionals.SkipRedownload.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChangeCategory.IsSet() {
		localVarQueryParams.Add("changeCategory", parameterToString(localVarOptionals.ChangeCategory.Value(), ""))
	}
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