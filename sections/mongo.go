package sections

import "github.com/teatak/config/v2"

type mongo struct {
	URI             string `yaml:"uri,omitempty"`
	Database        string `yaml:"database,omitempty"`
	ConnectTimeout  uint64 `yaml:"connectTimeout,omitempty"` //second
	MaxPoolSize     uint64 `yaml:"maxPoolSize,omitempty"`
	MinPoolSize     uint64 `yaml:"minPoolSize,omitempty"`
	MaxConnecting   uint64 `yaml:"maxConnecting,omitempty"`
	MaxConnIdleTime uint64 `yaml:"maxConnIdleTime,omitempty"` //second
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
