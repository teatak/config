package sections

import "github.com/teatak/config/v2"

type mongo struct {
	URI            string `yaml:"uri,omitempty"`
	Database       string `yaml:"database,omitempty"`
	ConnectTimeout uint64 `yaml:"connect_timeout,omitempty"` //second
	MaxPoolSize    uint64 `yaml:"max_pool_size,omitempty"`   //100
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
