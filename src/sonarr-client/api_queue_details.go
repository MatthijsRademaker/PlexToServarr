/*
 * Sonarr
 *
 * Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientSonarr

import (
	"context"
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

type QueueDetailsApiService service

/*
QueueDetailsApiService
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *QueueDetailsApiApiV3QueueDetailsGetOpts - Optional Parameters:
     * @param "SeriesId" (optional.Int32) -
     * @param "EpisodeIds" (optional.Interface of []int32) -
     * @param "IncludeSeries" (optional.Bool) -
     * @param "IncludeEpisode" (optional.Bool) -
@return []QueueResource
*/

type QueueDetailsApiApiV3QueueDetailsGetOpts struct {
	SeriesId       optional.Int32
	EpisodeIds     optional.Interface
	IncludeSeries  optional.Bool
	IncludeEpisode optional.Bool
}

func (a *QueueDetailsApiService) ApiV3QueueDetailsGet(ctx context.Context, localVarOptionals *QueueDetailsApiApiV3QueueDetailsGetOpts) ([]QueueResource, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []QueueResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/api/v3/queue/details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.SeriesId.IsSet() {
		localVarQueryParams.Add("seriesId", parameterToString(localVarOptionals.SeriesId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EpisodeIds.IsSet() {
		localVarQueryParams.Add("episodeIds", parameterToString(localVarOptionals.EpisodeIds.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeSeries.IsSet() {
		localVarQueryParams.Add("includeSeries", parameterToString(localVarOptionals.IncludeSeries.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IncludeEpisode.IsSet() {
		localVarQueryParams.Add("includeEpisode", parameterToString(localVarOptionals.IncludeEpisode.Value(), ""))
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
			var v []QueueResource
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
