package sections

import "github.com/teatak/config/v2"

type wechatpay struct {
	MchID      string `yaml:"mchID,omitempty"`
	Key        string `yaml:"key,omitempty"`
	SerialNo   string `yaml:"serialNo,omitempty"`
	PrivateKey string `yaml:"privateKey,omitempty"`
	PublicKey  string `yaml:"publicKey,omitempty"`
	NotifyUrl  string `yaml:"notifyUrl,omitempty"`
}

type wechatpaySection map[string]*wechatpay

func (s wechatpaySection) SectionName() string {
	return "wechatpay"
}

var WechatPay = wechatpaySection{}

func init() {
	config.Load(WechatPay)
}
