package main

import (
	"fmt"
	"os"

	"bookstore/api/internal/config"
	"bookstore/api/internal/handler"
	"bookstore/api/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := pwd + "/api/etc/bookstore-api.yaml"
	var c config.Config
	conf.MustLoad(configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
