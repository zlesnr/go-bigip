package bigip

const (
	uriSys            = "sys"
	uriSoftware       = "software"
	uriVolume         = "volume"
	uriHardware       = "hardware"
	uriGlobalSettings = "global-settings"
	//uriPlatform = "?$select=platform"
)

type Volumes struct {
	Volumes []Volume `json:"items,omitempty"`
}

type Volume struct {
	Name       string `json:"items,omitempty"`
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	Active     bool   `json:"active,omitempty"`
	BaseBuild  string `json:"basebuild,omitempty"`
	Build      string `json:"build,omitempty"`
	Product    string `json:"product,omitempty"`
	Status     string `json:"status,omitempty"`
	Version    string `json:"version,omitempty"`
}

// Volumes returns a list of Software Volumes.
func (b *BigIP) Volumes() (*Volumes, error) {
	var volumes Volumes
	err, _ := b.getForEntity(&volumes, uriSys, uriSoftware, uriVolume)
	if err != nil {
		return nil, err
	}

	return &volumes, nil
}

// type Hardware struct {
// 	Entries []HardwareEntry `json:"entries,omitempty"`
// }
//
// type HardwareEntry struct {
// }
