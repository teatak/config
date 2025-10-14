package sections

import "github.com/teatak/config/v2"

type redis struct {
	// 单个主机或集群配置
	// 例如：[]string{"192.168.1.10:6379"}
	Addrs []string `yaml:"addrs,omitempty"`
	// 哨兵 Master Name，仅适用于 `Failover Client`
	MasterName string `yaml:"masterName,omitempty"`
	// ClientName 和 `Options` 相同，会对每个Node节点的每个网络连接配置
	ClientName string `yaml:"clientName,omitempty"`
	// 设置 DB, 只针对 `Redis Client` 和 `Failover Client`
	DB               int    `yaml:"db,omitempty"`
	Username         string `yaml:"username,omitempty"`
	Password         string `yaml:"password,omitempty"`
	SentinelUsername string `yaml:"sentinelUsername,omitempty"`
	SentinelPassword string `yaml:"sentinelPassword,omitempty"`
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
