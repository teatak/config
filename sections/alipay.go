package sections

import "github.com/teatak/config"

type alipay struct {
	ClientId        string `yaml:"clientId,omitempty"`
	Gateway         string `yaml:"gateway,omitempty"`
	PrivateKey      string `yaml:"privateKey,omitempty"`
	AliPayPublicKey string `yaml:"aliPayPublicKey,omitempty"`
}

func (s *alipay) SectionName() string {
	return "alipay"
}

var Alipay = &alipay{}

func init() {
	config.Load(Alipay)
}
