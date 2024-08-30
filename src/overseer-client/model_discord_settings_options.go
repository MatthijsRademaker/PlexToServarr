/*
 * Overseerr API
 *
 * This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DiscordSettingsOptions struct {
	BotUsername string `json:"botUsername,omitempty"`
	BotAvatarUrl string `json:"botAvatarUrl,omitempty"`
	WebhookUrl string `json:"webhookUrl,omitempty"`
	EnableMentions bool `json:"enableMentions,omitempty"`
}
