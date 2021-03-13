package sections

import "github.com/teatak/config"

type server struct {
	Name string `yaml:"name,omitempty"`
	Port int    `yaml:"port,omitempty"`
}

func (s *server) SectionName() string {
	return "server"
}

var Server = &server{}

func init() {
	config.Load(Server)
}
