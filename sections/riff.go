package sections

import "github.com/teatak/config/v2"

type riff struct {
	Url string `yaml:"url,omitempty"`
}

var Riff = config.Register(&riff{})
