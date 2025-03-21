package main

import (
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/lance977/hmms-zero/rpc/notification/internal/config"
	"github.com/lance977/hmms-zero/rpc/notification/internal/server"
	"github.com/lance977/hmms-zero/rpc/notification/internal/svc"
	"github.com/lance977/hmms-zero/rpc/notification/notification"
)

var configFile = flag.String("f", "etc/notification.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		notification.RegisterNotificationServer(grpcServer, server.NewNotificationServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
