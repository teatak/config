package sections

import "github.com/teatak/config"

type gitlab struct {
	ClientId     string `yaml:"clientId,omitempty"`
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
