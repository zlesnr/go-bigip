package shared

// https://${F5_HOST}/mgmt/tm/shared/licensing/activation
type License struct {
	BaseRegKey            string `json:"baseRegKey,omitempty"`
	IsAutomaticActivation bool   `json:"isAutomaticActivation,omitempty"`
}
