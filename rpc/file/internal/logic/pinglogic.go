package logic

import (
	"context"
	"hmms-zero/rpc/file/file"
	"hmms-zero/rpc/file/internal/svc"

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

func (l *PingLogic) Ping(in *file.Request) (*file.Response, error) {
	// todo: add your logic here and delete this line

	return &file.Response{}, nil
}
