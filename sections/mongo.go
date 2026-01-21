package sections

import "github.com/teatak/config/v2"

type mongo struct {
	URI               string `yaml:"uri,omitempty"`
	Database          string `yaml:"database,omitempty"`
	MaxPoolSize       uint64 `yaml:"maxPoolSize,omitempty"`
	MinPoolSize       uint64 `yaml:"minPoolSize,omitempty"`
	ConnectTimeoutMS  uint64 `yaml:"connectTimeoutMS,omitempty"`  //ms
	MaxConnIdleTimeMS uint64 `yaml:"maxConnIdleTimeMS,omitempty"` //ms
	MaxConnecting     uint64 `yaml:"maxConnecting,omitempty"`
}

type mongoSection map[string]*mongo

func (s mongoSection) Default() *mongo {
	return s["default"]
}

var Mongo = mongoSection(config.RegisterMap[string, *mongo]("mongo"))
