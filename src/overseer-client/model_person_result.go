/*
 * Overseerr API
 *
 * This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type PersonResult struct {
	Id float64 `json:"id,omitempty"`
	ProfilePath string `json:"profilePath,omitempty"`
	Adult bool `json:"adult,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	KnownFor []OneOfPersonResultKnownForItems `json:"knownFor,omitempty"`
}
