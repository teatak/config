package sections

import "github.com/teatak/config/v2"

type alipay struct {
	AppID      string `yaml:"appID,omitempty"`
	Gateway    string `yaml:"gateway,omitempty"`
	PrivateKey string `yaml:"privateKey,omitempty"`
	PublicKey  string `yaml:"publicKey,omitempty"`
	NotifyUrl  string `yaml:"notifyUrl,omitempty"`
}

var Alipay = config.RegisterMap[*alipay]("alipay")
