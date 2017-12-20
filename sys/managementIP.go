package sys

type ManagementIP struct {
	Addresses []ManagementIPAddress
}

type ManagementIPAddress struct {
	Name       string `json:"items,omitempty"`
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
}
