package sections

import "github.com/teatak/config"

type smtp struct {
	Address  string `yaml:"address,omitempty"`
	Name     string `yaml:"name,omitempty"`
	UserName string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

func (s *smtp) SectionName() string {
	return "smtp"
}

var Smtp = &smtp{}

func init() {
	config.Load(Smtp)
}
