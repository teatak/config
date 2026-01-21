package sections

import "github.com/teatak/config/v2"

type log struct {
	Handler string `yaml:"handler,omitempty"`
	Level   string `yaml:"level,omitempty"`
}

var Log = config.Register(&log{})
