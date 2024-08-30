/*
 * Overseerr API
 *
 * This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RadarrSettings struct {
	Id float64 `json:"id,omitempty"`
	Name string `json:"name"`
	Hostname string `json:"hostname"`
	Port float64 `json:"port"`
	ApiKey string `json:"apiKey"`
	UseSsl bool `json:"useSsl"`
	BaseUrl string `json:"baseUrl,omitempty"`
	ActiveProfileId float64 `json:"activeProfileId"`
	ActiveProfileName string `json:"activeProfileName"`
	ActiveDirectory string `json:"activeDirectory"`
	Is4k bool `json:"is4k"`
	MinimumAvailability string `json:"minimumAvailability"`
	IsDefault bool `json:"isDefault"`
	ExternalUrl string `json:"externalUrl,omitempty"`
	SyncEnabled bool `json:"syncEnabled,omitempty"`
	PreventSearch bool `json:"preventSearch,omitempty"`
}
