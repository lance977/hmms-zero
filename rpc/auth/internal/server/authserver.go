// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: auth.proto

package server

import (
	"context"

	"github.com/lance977/hmms-zero/rpc/auth/auth"
	"github.com/lance977/hmms-zero/rpc/auth/internal/logic"
	"github.com/lance977/hmms-zero/rpc/auth/internal/svc"
)

type AuthServer struct {
	svcCtx *svc.ServiceContext
	auth.UnimplementedAuthServer
}

func NewAuthServer(svcCtx *svc.ServiceContext) *AuthServer {
	return &AuthServer{
		svcCtx: svcCtx,
	}
}

func (s *AuthServer) Ping(ctx context.Context, in *auth.Request) (*auth.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
