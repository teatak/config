package sections

import "github.com/teatak/config"

type mongo struct {
	Uri            string   `yaml:"uri,omitempty"`
	Hosts          []string `yaml:"hosts,omitempty"`
	Database       string   `yaml:"database,omitempty"`
	ReplicaSetName string   `yaml:"replica_set_name,omitempty"`
	Username       string   `yaml:"username,omitempty"`
	Password       string   `yaml:"password,omitempty"`
	Source         string   `yaml:"source,omitempty"`
	PoolSize       int      `yaml:"pool_size,omitempty"`
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
