package main

import (
	"flag"
	"fmt"
	"github.com/lance977/hmms-zero/rpc/home/home"
	"github.com/lance977/hmms-zero/rpc/home/internal/config"
	"github.com/lance977/hmms-zero/rpc/home/internal/server"
	"github.com/lance977/hmms-zero/rpc/home/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/home.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		home.RegisterHomeServer(grpcServer, server.NewHomeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
