package sections

import "github.com/teatak/config/v2"

type nacos struct {
	Enable      bool   `yaml:"enable"`
	IpAddr      string `yaml:"ipAddr"`
	Port        uint64 `yaml:"port"`
	NamespaceId string `yaml:"namespaceId"`
	DataId      string `yaml:"dataId"`
	Group       string `yaml:"group"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	Mode        string `yaml:"mode"` // merge or overwrite
}

var Nacos = config.Register(&nacos{})
