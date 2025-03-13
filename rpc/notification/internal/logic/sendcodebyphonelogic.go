package logic

import (
	"context"

	"github.com/lance977/hmms-zero/rpc/notification/internal/svc"
	"github.com/lance977/hmms-zero/rpc/notification/notification"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendCodeByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendCodeByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendCodeByPhoneLogic {
	return &SendCodeByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发送手机验证码
func (l *SendCodeByPhoneLogic) SendCodeByPhone(in *notification.SendCodeByPhoneReq) (*notification.SendCodeByPhoneResp, error) {
	// todo: add your logic here and delete this line

	return &notification.SendCodeByPhoneResp{}, nil
}
