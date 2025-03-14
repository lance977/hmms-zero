package logic

import (
	"context"

	"hmms-zero/rpc/user/internal/svc"
	"hmms-zero/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *user.PingReq) (*user.PingResp, error) {
	// todo: add your logic here and delete this line

	return &user.PingResp{}, nil
}
