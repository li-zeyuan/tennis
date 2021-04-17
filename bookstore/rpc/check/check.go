package main

import (
	"fmt"
	"os"

	"bookstore/rpc/check/check"
	"bookstore/rpc/check/internal/config"
	"bookstore/rpc/check/internal/server"
	"bookstore/rpc/check/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)


func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := pwd + "/rpc/check/etc/check.yaml"

	var c config.Config
	conf.MustLoad(configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCheckServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		check.RegisterCheckServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
