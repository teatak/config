package sections

import "github.com/teatak/config/v2"

type smtp struct {
	Address  string `yaml:"address,omitempty"`
	Name     string `yaml:"name,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

var Smtp = config.Register(&smtp{})
