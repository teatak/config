package sections

import "github.com/teatak/config"

type wechat struct {
	AppID     string `yaml:"appID,omitempty"`
	AppSecret string `yaml:"appSecret,omitempty"`
}

type wechatSection map[string]*wechat

func (s wechatSection) SectionName() string {
	return "wechat"
}

var Wechat = wechatSection{}

func init() {
	config.Load(Wechat)
}
