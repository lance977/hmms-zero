package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/lance977/hmms-zero/rpc/notification/internal/svc"
	"github.com/lance977/hmms-zero/rpc/notification/notification"
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

func (l *PingLogic) Ping(in *notification.PingReq) (*notification.PingResp, error) {
	return &notification.PingResp{Pong: "pong"}, nil
}
