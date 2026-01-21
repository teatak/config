package sections

import "github.com/teatak/config/v2"

type aliyun struct {
	AccessKeyID  string `yaml:"accessKeyID,omitempty"`
	AccessSecret string `yaml:"accessSecret,omitempty"`
}

type aliyunSection map[string]*aliyun

func (s aliyunSection) Default() *aliyun {
	return s["default"]
}

var Aliyun = aliyunSection(config.RegisterMap[string, *aliyun]("aliyun"))
