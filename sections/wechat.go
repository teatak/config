package sections

import "github.com/teatak/config/v2"

type wechat struct {
	AppID     string `yaml:"appID,omitempty"`
	AppSecret string `yaml:"appSecret,omitempty"`
}

var Wechat = config.RegisterMap[*wechat]("wechat")
