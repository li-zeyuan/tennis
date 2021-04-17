package main

import (
	"fmt"
	"os"

	"bookstore/rpc/add/add"
	"bookstore/rpc/add/internal/config"
	"bookstore/rpc/add/internal/server"
	"bookstore/rpc/add/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := pwd + "/rpc/add/etc/add.yaml"

	var c config.Config
	conf.MustLoad(configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewAdderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		add.RegisterAdderServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
