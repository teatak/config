package sections

import "github.com/teatak/config/v2"

type github struct {
	ClientID     string `yaml:"clientID,omitempty"`
	ClientSecret string `yaml:"clientSecret,omitempty"`
}

var Github = config.Register(&github{})
