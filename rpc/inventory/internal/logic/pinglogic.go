package logic

import (
	"context"
	"hmms-zero/rpc/inventory/internal/svc"
	"hmms-zero/rpc/inventory/inventory"

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

func (l *PingLogic) Ping(in *inventory.Request) (*inventory.Response, error) {
	// todo: add your logic here and delete this line

	return &inventory.Response{}, nil
}
