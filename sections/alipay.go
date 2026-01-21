package sections

import "github.com/teatak/config/v2"

type alipay struct {
	AppID      string `yaml:"appID,omitempty"`
	Gateway    string `yaml:"gateway,omitempty"`
	PrivateKey string `yaml:"privateKey,omitempty"`
	PublicKey  string `yaml:"publicKey,omitempty"`
	NotifyUrl  string `yaml:"notifyUrl,omitempty"`
}

type alipaySection map[string]*alipay

func (s alipaySection) Default() *alipay {
	return s["default"]
}

var Alipay = alipaySection(config.RegisterMap[string, *alipay]("alipay"))
