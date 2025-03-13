package logic

import (
	"context"

	"github.com/lance977/hmms-zero/rpc/notification/internal/svc"
	"github.com/lance977/hmms-zero/rpc/notification/notification"

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

func (l *PingLogic) Ping(in *notification.PingReq) (*notification.PingResp, error) {
	// todo: add your logic here and delete this line

	return &notification.PingResp{}, nil
}
