package sections

import "github.com/teatak/config"

type github struct {
	ClientId     string `yaml:"clientId,omitempty"`
	ClientSecret string `yaml:"clientSecret,omitempty"`
}

func (s *github) SectionName() string {
	return "github"
}

var Github = &github{}

func init() {
	config.Load(Github)
}
