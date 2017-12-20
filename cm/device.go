package cm

type Devices struct {
	Devices []Device `json:"items,omitempty"`
}

// /cm/device
type Device struct {
	Kind          string   `json:"kind"`
	Name          string   `json:"name"`
	Partition     string   `json:"partition"`
	FullPath      string   `json:"fullPath"`
	Generation    int      `json:"generation"`
	SelfLink      string   `json:"selfLink"`
	ActiveModules []string `json:"activeModules"`
	BaseMac       string   `json:"baseMac"`
	Build         string   `json:"build"`
	Cert          string   `json:"cert"`
	CertReference struct {
		Link string `json:"link"`
	} `json:"certReference"`
	ChassisID     string `json:"chassisId"`
	ChassisType   string `json:"chassisType"`
	ConfigsyncIP  string `json:"configsyncIp"`
	Edition       string `json:"edition"`
	FailoverState string `json:"failoverState"`
	HaCapacity    int    `json:"haCapacity"`
	Hostname      string `json:"hostname"`
	Key           string `json:"key"`
	KeyReference  struct {
		Link string `json:"link"`
	} `json:"keyReference"`
	ManagementIP      string   `json:"managementIp"`
	MarketingName     string   `json:"marketingName"`
	MirrorIP          string   `json:"mirrorIp"`
	MirrorSecondaryIP string   `json:"mirrorSecondaryIp"`
	MulticastIP       string   `json:"multicastIp"`
	MulticastPort     int      `json:"multicastPort"`
	OptionalModules   []string `json:"optionalModules"`
	PlatformID        string   `json:"platformId"`
	Product           string   `json:"product"`
	SelfDevice        string   `json:"selfDevice"`
	TimeZone          string   `json:"timeZone"`
	Version           string   `json:"version"`
}
