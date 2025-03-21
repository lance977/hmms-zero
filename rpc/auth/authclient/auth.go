// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: auth.proto

package authclient

import (
	"context"

	"github.com/lance977/hmms-zero/rpc/auth/auth"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = auth.Request
	Response = auth.Response

	Auth interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultAuth struct {
		cli zrpc.Client
	}
)

func NewAuth(cli zrpc.Client) Auth {
	return &defaultAuth{
		cli: cli,
	}
}

func (m *defaultAuth) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := auth.NewAuthClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
