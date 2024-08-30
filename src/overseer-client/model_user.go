/*
 * Overseerr API
 *
 * This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type User struct {
	Id int32 `json:"id"`
	Email string `json:"email"`
	Username string `json:"username,omitempty"`
	PlexToken string `json:"plexToken,omitempty"`
	PlexUsername string `json:"plexUsername,omitempty"`
	UserType int32 `json:"userType,omitempty"`
	Permissions float64 `json:"permissions,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	RequestCount float64 `json:"requestCount,omitempty"`
}
