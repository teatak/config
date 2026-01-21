package sections

import "github.com/teatak/config/v2"

type aliyun struct {
	AccessKeyID  string `yaml:"accessKeyID,omitempty"`
	AccessSecret string `yaml:"accessSecret,omitempty"`
}

var Aliyun = config.RegisterMap[*aliyun]("aliyun")
