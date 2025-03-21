// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: notification.proto

package notificationclient

import (
	"context"

	"github.com/lance977/hmms-zero/rpc/notification/notification"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PingReq             = notification.PingReq
	PingResp            = notification.PingResp
	SendCodeByPhoneReq  = notification.SendCodeByPhoneReq
	SendCodeByPhoneResp = notification.SendCodeByPhoneResp

	Notification interface {
		Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
		// 发送手机验证码
		SendCodeByPhone(ctx context.Context, in *SendCodeByPhoneReq, opts ...grpc.CallOption) (*SendCodeByPhoneResp, error)
	}

	defaultNotification struct {
		cli zrpc.Client
	}
)

func NewNotification(cli zrpc.Client) Notification {
	return &defaultNotification{
		cli: cli,
	}
}

func (m *defaultNotification) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	client := notification.NewNotificationClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

// 发送手机验证码
func (m *defaultNotification) SendCodeByPhone(ctx context.Context, in *SendCodeByPhoneReq, opts ...grpc.CallOption) (*SendCodeByPhoneResp, error) {
	client := notification.NewNotificationClient(m.cli.Conn())
	return client.SendCodeByPhone(ctx, in, opts...)
}
