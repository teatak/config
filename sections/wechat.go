package sections

import "github.com/teatak/config/v2"

type wechat struct {
	AppID     string `yaml:"appID,omitempty"`
	AppSecret string `yaml:"appSecret,omitempty"`
}

type wechatSection map[string]*wechat

func (s wechatSection) Default() *wechat {
	return s["default"]
}

var Wechat = wechatSection(config.RegisterMap[string, *wechat]("wechat"))
