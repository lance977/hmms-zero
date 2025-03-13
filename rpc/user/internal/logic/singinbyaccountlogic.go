package logic

import (
	"context"

	"hmms-zero/rpc/user/internal/svc"
	"hmms-zero/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type SingInByAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSingInByAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SingInByAccountLogic {
	return &SingInByAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SingInByAccountLogic) SingInByAccount(in *user.SignInByAccountReq) (*user.SignInResp, error) {
	// todo: add your logic here and delete this line

	return &user.SignInResp{}, nil
}
