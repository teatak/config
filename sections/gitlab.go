package sections

import "github.com/teatak/config/v2"

type gitlab struct {
	ClientID     string `yaml:"clientID,omitempty"`
	ClientSecret string `yaml:"clientSecret,omitempty"`
	RedirectUri  string `yaml:"redirectUri,omitempty"`
}

func (s *gitlab) SectionName() string {
	return "gitlab"
}

var Gitlab = &gitlab{}

func init() {
	config.Load(Gitlab)
}
