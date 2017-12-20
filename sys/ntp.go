package sys

type NTP struct {
	Kind        string   `json:"kind,omitempty"`
	SelfLink    string   `json:"selfLink,omitempty"`
	Description string   `json:"description,omitempty"`
	Servers     []string `json:"servers,omitempty"`
	Timezone    string   `json:"timezone,omitempty"`
}
