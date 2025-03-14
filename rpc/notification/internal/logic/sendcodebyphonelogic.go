package logic

import (
	"context"
	"errors"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lance977/hmms-zero/rpc/notification/internal/svc"
	"github.com/lance977/hmms-zero/rpc/notification/notification"
)

const CodeOK = "OK"

type SendCodeByPhoneLogic struct {
	ctx       context.Context
	svcCtx    *svc.ServiceContext
	smsClient *dysmsapi.Client
	logx.Logger
}

func NewSendCodeByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeByPhoneLogic {
	client, err := NewSmsClient(svcCtx)
	if err != nil {
		panic(err)
	}
	return &SendCodeByPhoneLogic{
		ctx:       ctx,
		svcCtx:    svcCtx,
		smsClient: client,
		Logger:    logx.WithContext(ctx),
	}
}

func NewSmsClient(svcCtx *svc.ServiceContext) (*dysmsapi.Client, error) {
	if svcCtx == nil || svcCtx.Config.Sms == nil {
		return nil, errors.New("config is nil")
	}
	smsCfg := svcCtx.Config.Sms
	return dysmsapi.NewClient(&openapi.Config{
		AccessKeyId:     &smsCfg.AccessKeyID,
		AccessKeySecret: &smsCfg.AccessKeySecret,
	})
}

// SendCodeByPhone 发送手机验证码
func (l *SendCodeByPhoneLogic) SendCodeByPhone(in *notification.SendCodeByPhoneReq) (*notification.SendCodeByPhoneResp, error) {
	code := `{"code":"` + in.Code + `"}`
	req := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  &in.PhoneNumber,
		SignName:      &l.svcCtx.Config.Sms.SignName,
		TemplateCode:  &l.svcCtx.Config.Sms.TemplateCode,
		TemplateParam: &code,
	}
	resp, err := l.smsClient.SendSms(req)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}

	if *resp.Body.Code != CodeOK {
		l.Logger.Errorf("SendSms error: code=%s, message=%s, requestId=%s",
			*resp.Body.Code, *resp.Body.Message, *resp.Body.RequestId)
		return nil, errors.New(*resp.Body.Message)
	}
	return &notification.SendCodeByPhoneResp{Success: true}, nil
}
