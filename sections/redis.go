package sections

import "github.com/teatak/config"

type redis struct {
	Sentinel *struct {
		Master    string   `yaml:"master,omitempty"`
		Addresses []string `yaml:"addresses,omitempty"`
	} `yaml:"sentinel,omitempty"`
	Cluster *struct {
		Addresses []string `yaml:"addresses,omitempty"`
	} `yaml:"cluster,omitempty"`
	Address  string `yaml:"address,omitempty"`
	Password string `yaml:"password,omitempty"`
	Db       int    `yaml:"db,omitempty"`
}

type redisSection map[string]*redis

func (s *redisSection) SectionName() string {
	return "redis"
}

func (s *redisSection) Default() *redis {
	return Redis["default"]
}

var Redis = redisSection{}

func init() {
	config.Load(&Redis)
}
