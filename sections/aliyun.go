package sections

import "github.com/teatak/config/v2"

type aliyun struct {
	AccessKeyID        string `yaml:"accessKeyID,omitempty"`
	AccessSecret       string `yaml:"accessSecret,omitempty"`
	SignName           string `yaml:"signName,omitempty"`           // 短信签名（需在阿里云控制台申请）
	VerifyCodeTemplate string `yaml:"verifyCodeTemplate,omitempty"` // 验证码短信模板编号（如 SMS_123456789）
}

var Aliyun = config.RegisterMap[*aliyun]("aliyun")
