package config

import (
	"errors"

	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Sms *SmsConfig
}

type SmsConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	SignName        string // 签名
	TemplateCode    string // 短信模板
}

func (c SmsConfig) Validate() error {
	if len(c.AccessKeyID) == 0 {
		return errors.New("SMS AccessKeyID 不能为空")
	}
	if len(c.AccessKeySecret) == 0 {
		return errors.New("SMS AccessKeySecret 不能为空")
	}
	if len(c.SignName) == 0 {
		return errors.New("SMS SignName 不能为空")
	}
	if len(c.TemplateCode) == 0 {
		return errors.New("SMS TemplateCode 不能为空")
	}
	return nil
}
