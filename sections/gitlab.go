package sections

import "github.com/teatak/config/v2"

type gitlab struct {
	ClientID     string `yaml:"clientID,omitempty"`
	ClientSecret string `yaml:"clientSecret,omitempty"`
	RedirectUri  string `yaml:"redirectUri,omitempty"`
}

var Gitlab = config.Register(&gitlab{})
