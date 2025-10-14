package sections

import "github.com/teatak/config/v2"

type riff struct {
	Url string `yaml:"url,omitempty"`
}

func (s *riff) SectionName() string {
	return "riff"
}

var Riff = &riff{}

func init() {
	config.Load(Riff)
}
