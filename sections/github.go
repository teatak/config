package sections

import "github.com/teatak/config"

type github struct {
	ClientID     string `yaml:"clientID,omitempty"`
	ClientSecret string `yaml:"clientSecret,omitempty"`
}

func (s *github) SectionName() string {
	return "github"
}

var Github = &github{}

func init() {
	config.Load(Github)
}
