/*
 * Sonarr
 *
 * Sonarr API docs - The v3 API docs apply to both v3 and v4 versions of Sonarr. Some functionality may only be available in v4 of the Sonarr application.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swaggerClientSonarr

type ImportListResource struct {
	Id                       int32                `json:"id,omitempty"`
	Name                     string               `json:"name,omitempty"`
	Fields                   []Field              `json:"fields,omitempty"`
	ImplementationName       string               `json:"implementationName,omitempty"`
	Implementation           string               `json:"implementation,omitempty"`
	ConfigContract           string               `json:"configContract,omitempty"`
	InfoLink                 string               `json:"infoLink,omitempty"`
	Message                  *ProviderMessage     `json:"message,omitempty"`
	Tags                     []int32              `json:"tags,omitempty"`
	Presets                  []ImportListResource `json:"presets,omitempty"`
	EnableAutomaticAdd       bool                 `json:"enableAutomaticAdd,omitempty"`
	SearchForMissingEpisodes bool                 `json:"searchForMissingEpisodes,omitempty"`
	ShouldMonitor            *MonitorTypes        `json:"shouldMonitor,omitempty"`
	MonitorNewItems          *NewItemMonitorTypes `json:"monitorNewItems,omitempty"`
	RootFolderPath           string               `json:"rootFolderPath,omitempty"`
	QualityProfileId         int32                `json:"qualityProfileId,omitempty"`
	SeriesType               *SeriesTypes         `json:"seriesType,omitempty"`
	SeasonFolder             bool                 `json:"seasonFolder,omitempty"`
	ListType                 *ImportListType      `json:"listType,omitempty"`
	ListOrder                int32                `json:"listOrder,omitempty"`
	MinRefreshInterval       string               `json:"minRefreshInterval,omitempty"`
}
