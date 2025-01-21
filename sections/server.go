package sections

import "github.com/teatak/config"

type server struct {
	Environment  string   `yaml:"environment,omitempty"`
	Url          string   `yaml:"url,omitempty"`
	ShortUrl     string   `yaml:"shortUrl,omitempty"`
	AllowOrigins []string `yaml:"allowOrigins,omitempty"`
	Name         string   `yaml:"name,omitempty"`
	Port         int      `yaml:"port,omitempty"`
}

func (s *server) SectionName() string {
	return "server"
}

var Server = &server{}

func init() {
	config.Load(Server)
}
