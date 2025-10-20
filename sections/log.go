package sections

import "github.com/teatak/config/v2"

type log struct {
	Handler  string `yaml:"handler,omitempty"`
	LogLevel string `yaml:"logLevel,omitempty"`
}

func (s *log) SectionName() string {
	return "log"
}

var Log = &log{}

func init() {
	config.Load(Log)
}
