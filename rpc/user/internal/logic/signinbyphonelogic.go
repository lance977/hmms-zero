package logic

import (
	"context"

	"hmms-zero/rpc/user/internal/svc"
	"hmms-zero/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInByPhoneLogic {
	return &SignInByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInByPhoneLogic) SignInByPhone(in *user.SignInByPhoneReq) (*user.SignInResp, error) {
	// todo: add your logic here and delete this line

	return &user.SignInResp{}, nil
}
