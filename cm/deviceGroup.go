package cm

type DeviceGroups struct {
	DeviceGroups []DeviceGroup `json:"items,omitempty"`
}

type DeviceGroup struct {
	Kind                         string `json:"kind,omitempty"`
	Name                         string `json:"name,omitempty"`
	Partition                    string `json:"partition,omitempty"`
	FullPath                     string `json:"fullPath,omitempty"`
	Generation                   int    `json:"generation,omitempty"`
	SelfLink                     string `json:"selfLink,omitempty"`
	AsmSync                      string `json:"asmSync,omitempty"`
	AutoSync                     string `json:"autoSync,omitempty"`
	FullLoadOnSync               string `json:"fullLoadOnSync,omitempty"`
	IncrementalConfigSyncSizeMax int    `json:"incrementalConfigSyncSizeMax,omitempty"`
	NetworkFailover              string `json:"networkFailover,omitempty"`
	SaveOnAutoSync               string `json:"saveOnAutoSync,omitempty"`
	Type                         string `json:"type,omitempty"`
	DevicesReference             struct {
		Link            string `json:"link"`
		IsSubcollection bool   `json:"isSubcollection"`
	} `json:"devicesReference,omitempty"`
}
