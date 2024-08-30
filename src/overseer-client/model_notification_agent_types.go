/*
 * Overseerr API
 *
 * This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NotificationAgentTypes struct {
	Discord float64 `json:"discord,omitempty"`
	Email float64 `json:"email,omitempty"`
	Pushbullet float64 `json:"pushbullet,omitempty"`
	Pushover float64 `json:"pushover,omitempty"`
	Slack float64 `json:"slack,omitempty"`
	Telegram float64 `json:"telegram,omitempty"`
	Webhook float64 `json:"webhook,omitempty"`
	Webpush float64 `json:"webpush,omitempty"`
}