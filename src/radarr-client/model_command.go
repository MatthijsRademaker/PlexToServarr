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
	"time"
)

type Command struct {
	SendUpdatesToClient bool            `json:"sendUpdatesToClient,omitempty"`
	UpdateScheduledTask bool            `json:"updateScheduledTask,omitempty"`
	CompletionMessage   string          `json:"completionMessage,omitempty"`
	RequiresDiskAccess  bool            `json:"requiresDiskAccess,omitempty"`
	IsExclusive         bool            `json:"isExclusive,omitempty"`
	IsTypeExclusive     bool            `json:"isTypeExclusive,omitempty"`
	IsLongRunning       bool            `json:"isLongRunning,omitempty"`
	Name                string          `json:"name,omitempty"`
	LastExecutionTime   time.Time       `json:"lastExecutionTime,omitempty"`
	LastStartTime       time.Time       `json:"lastStartTime,omitempty"`
	Trigger             *CommandTrigger `json:"trigger,omitempty"`
	SuppressMessages    bool            `json:"suppressMessages,omitempty"`
	ClientUserAgent     string          `json:"clientUserAgent,omitempty"`
}