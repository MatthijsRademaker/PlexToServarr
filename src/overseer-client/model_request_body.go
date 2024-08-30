/*
 * Overseerr API
 *
 * This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RequestBody struct {
	MediaType         string    `json:"mediaType"`
	MediaId           float64   `json:"mediaId"`
	TvdbId            float64   `json:"tvdbId,omitempty"`
	Seasons           []float64 `json:"seasons,omitempty"`
	Is4k              bool      `json:"is4k,omitempty"`
	ServerId          float64   `json:"serverId,omitempty"`
	ProfileId         float64   `json:"profileId,omitempty"`
	RootFolder        string    `json:"rootFolder,omitempty"`
	LanguageProfileId float64   `json:"languageProfileId,omitempty"`
	UserId            float64   `json:"userId,omitempty"`
}
