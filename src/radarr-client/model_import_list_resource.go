/*
 * Radarr
 *
 * Radarr API docs
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientRadarr

type ImportListResource struct {
	Id                  int32                `json:"id,omitempty"`
	Name                string               `json:"name,omitempty"`
	Fields              []Field              `json:"fields,omitempty"`
	ImplementationName  string               `json:"implementationName,omitempty"`
	Implementation      string               `json:"implementation,omitempty"`
	ConfigContract      string               `json:"configContract,omitempty"`
	InfoLink            string               `json:"infoLink,omitempty"`
	Message             *ProviderMessage     `json:"message,omitempty"`
	Tags                []int32              `json:"tags,omitempty"`
	Presets             []ImportListResource `json:"presets,omitempty"`
	Enabled             bool                 `json:"enabled,omitempty"`
	EnableAuto          bool                 `json:"enableAuto,omitempty"`
	Monitor             *MonitorTypes        `json:"monitor,omitempty"`
	RootFolderPath      string               `json:"rootFolderPath,omitempty"`
	QualityProfileId    int32                `json:"qualityProfileId,omitempty"`
	SearchOnAdd         bool                 `json:"searchOnAdd,omitempty"`
	MinimumAvailability *MovieStatusType     `json:"minimumAvailability,omitempty"`
	ListType            *ImportListType      `json:"listType,omitempty"`
	ListOrder           int32                `json:"listOrder,omitempty"`
	MinRefreshInterval  string               `json:"minRefreshInterval,omitempty"`
}