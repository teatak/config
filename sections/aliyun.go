package sections

import "github.com/teatak/config"

type aliyun struct {
	AccessKeyID  string `yaml:"accessKeyID,omitempty"`
	AccessSecret string `yaml:"accessSecret,omitempty"`
}

type aliyunSection map[string]*aliyun

func (s *aliyunSection) SectionName() string {
	return "aliyun"
}

func (s *aliyunSection) Default() *aliyun {
	return Aliyun["default"]
}

var Aliyun = aliyunSection{}

func init() {
	config.Load(&Aliyun)
}
