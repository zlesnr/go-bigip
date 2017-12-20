package sys

type GlobalSettings struct {
	Kind                            string `json:"kind,omitempty"`
	SelfLink                        string `json:"selfLink,omitempty"`
	AwsAPIMaxConcurrency            int    `json:"awsApiMaxConcurrency,omitempty"`
	ConsoleInactivityTimeout        int    `json:"consoleInactivityTimeout,omitempty"`
	CustomAddr                      string `json:"customAddr,omitempty"`
	FailsafeAction                  string `json:"failsafeAction,omitempty"`
	FileBlacklistPathPrefix         string `json:"fileBlacklistPathPrefix,omitempty"`
	FileBlacklistReadOnlyPathPrefix string `json:"fileBlacklistReadOnlyPathPrefix,omitempty"`
	FileLocalPathPrefix             string `json:"fileLocalPathPrefix,omitempty"`
	FileWhitelistPathPrefix         string `json:"fileWhitelistPathPrefix,omitempty"`
	GuiSecurityBanner               string `json:"guiSecurityBanner,omitempty"`
	GuiSecurityBannerText           string `json:"guiSecurityBannerText,omitempty"`
	GuiSetup                        string `json:"guiSetup,omitempty"`
	HostAddrMode                    string `json:"hostAddrMode,omitempty"`
	Hostname                        string `json:"hostname,omitempty"`
	LcdDisplay                      string `json:"lcdDisplay,omitempty"`
	LedLocator                      string `json:"ledLocator,omitempty"`
	MgmtDhcp                        string `json:"mgmtDhcp,omitempty"`
	NetReboot                       string `json:"netReboot,omitempty"`
	PasswordPrompt                  string `json:"passwordPrompt,omitempty"`
	QuietBoot                       string `json:"quietBoot,omitempty"`
	UsernamePrompt                  string `json:"usernamePrompt,omitempty"`
}

// Volumes returns a list of Software Volumes.
func (b *BigIP) GlobalSettings() (*GlobalSettings, error) {
	var globalSettings GlobalSettings
	err, _ := b.getForEntity(&globalSettings, uriSys, uriGlobalSettings)
	if err != nil {
		return nil, err
	}

	return &globalSettings, nil
}

func (b *BigIP) ModifyGlobalSettings(config *GlobalSettings) error {
	return b.put(config, uriSys, uriGlobalSettings)
}
