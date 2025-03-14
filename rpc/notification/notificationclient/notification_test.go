package notificationclient

import (
	"context"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
)

func newClient() (Notification, error) {
	clientConf := zrpc.RpcClientConf{}
	err := conf.FillDefault(&clientConf)
	if err != nil {
		return nil, err
	}
	clientConf.Timeout = 20000 // 超时时间
	clientConf.Endpoints = []string{"127.0.0.1:8080"}
	conn := zrpc.MustNewClient(clientConf)
	client := NewNotification(conn)
	return client, nil
}

func TestDefaultNotification_Ping(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.Ping(context.TODO(), &PingReq{Ping: "ping"})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Success: %+v\n", resp)
}

func TestDefaultNotification_SendCodeByPhone(t *testing.T) {
	client, err := newClient()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.SendCodeByPhone(context.TODO(), &SendCodeByPhoneReq{
		PhoneNumber: "15670761927",
		Code:        "302943",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Success: %+v\n", resp)
}
