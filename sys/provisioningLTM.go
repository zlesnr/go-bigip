package sys

type ProvisionLTM struct {
	Name        string `json:"items,omitempty"`
	FullPath    string `json:"fullPath,omitempty"`
	Generation  int    `json:"generation,omitempty"`
	SelfLink    string `json:"selfLink,omitempty"`
	CPURatio    int    `json:"cpuRatio,omitempty"`
	DiskRatio   int    `json:"diskRatio,omitempty"`
	Level       string `json:"level,omitempty"`
	MemoryRatio int    `json:"memoryRatio,omitempty"`
}
