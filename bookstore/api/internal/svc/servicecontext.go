package svc

import (
	"bookstore/api/internal/config"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/check/checker"
	"github.com/tal-tech/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder
	Checker checker.Check
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewCheck(zrpc.MustNewClient(c.Check)),
	}
}
