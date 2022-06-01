package sections

import "github.com/teatak/config"

type wechat struct {
	AppId     string `yaml:"appId,omitempty"`
	AppSecret string `yaml:"appSecret,omitempty"`
	Type      string `yaml:"type,omitempty"`
}

type wechatSection map[string]*wechat

func (s wechatSection) SectionName() string {
	return "wechat"
}

var Wechat = wechatSection{}

func init() {
	config.Load(Wechat)
}
