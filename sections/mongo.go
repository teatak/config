package sections

import "github.com/teatak/config/v2"

type mongo struct {
	URI      string `yaml:"uri,omitempty"`
	Database string `yaml:"database,omitempty"`
}

type mongoSection map[string]*mongo

func (s *mongoSection) SectionName() string {
	return "mongo"
}

func (s *mongoSection) Default() *mongo {
	return Mongo["default"]
}

var Mongo = mongoSection{}

func init() {
	config.Load(&Mongo)
}
