package sys

type DNS struct {
	Kind         string   `json:"kind,omitempty"`
	SelfLink     string   `json:"selfLink,omitempty"`
	Description  string   `json:"description,omitempty"`
	NameServerrs []string `json:"nameServers,omitempty"`
	NumberOfDots int      `json:"numberOfDots,omitempty"`
	Search       []string `json:"search,omitempty"`
}
