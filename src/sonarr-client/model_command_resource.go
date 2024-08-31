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
	"time"
)

type CommandResource struct {
	Id                  int32            `json:"id,omitempty"`
	Name                string           `json:"name,omitempty"`
	CommandName         string           `json:"commandName,omitempty"`
	Message             string           `json:"message,omitempty"`
	Body                *Command         `json:"body,omitempty"`
	Priority            *CommandPriority `json:"priority,omitempty"`
	Status              *CommandStatus   `json:"status,omitempty"`
	Result              *CommandResult   `json:"result,omitempty"`
	Queued              time.Time        `json:"queued,omitempty"`
	Started             time.Time        `json:"started,omitempty"`
	Ended               time.Time        `json:"ended,omitempty"`
	Duration            string           `json:"duration,omitempty"`
	Exception           string           `json:"exception,omitempty"`
	Trigger             *CommandTrigger  `json:"trigger,omitempty"`
	ClientUserAgent     string           `json:"clientUserAgent,omitempty"`
	StateChangeTime     time.Time        `json:"stateChangeTime,omitempty"`
	SendUpdatesToClient bool             `json:"sendUpdatesToClient,omitempty"`
	UpdateScheduledTask bool             `json:"updateScheduledTask,omitempty"`
	LastExecutionTime   time.Time        `json:"lastExecutionTime,omitempty"`
}