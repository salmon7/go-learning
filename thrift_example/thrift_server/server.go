package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/salmon7/go-learning/thrift_example/gen-go/echo"
)

type EchoServerImpl struct {
}

func (p *EchoServerImpl) Echo(ctx context.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
	fmt.Printf("message from client: %v\n", req.GetMsg())
	res := &echo.EchoRes{
		Msg: req.GetMsg(),
	}
	return res, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":3000")
	if err != nil {
		panic(err)
	}
	processor := echo.NewEchoProcessor(&EchoServerImpl{})
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		thrift.NewTBufferedTransportFactory(8192),
		thrift.NewTCompactProtocolFactory(),
	)
	if err := server.Serve(); err != nil {
		panic(err)
	}
}
