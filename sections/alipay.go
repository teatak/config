package sections

import "github.com/teatak/config"

type alipay struct {
	AppID      string `yaml:"appID,omitempty"`
	Gateway    string `yaml:"gateway,omitempty"`
	PrivateKey string `yaml:"privateKey,omitempty"`
	PublicKey  string `yaml:"publicKey,omitempty"`
	NotifyUrl  string `yaml:"notifyUrl,omitempty"`
}

type alipaySection map[string]*alipay

func (s alipaySection) SectionName() string {
	return "alipay"
}

var AliPay = alipaySection{}

func init() {
	config.Load(AliPay)
}
