package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/salmon7/go-learning/thrift_example/gen-go/echo"
)

func main() {
	var transport thrift.TTransport
	var err error

	transport, err = thrift.NewTSocket("localhost:3000")
	if err != nil {
		fmt.Printf("NewTSocket failed. err: [%v]\n", err)
		return
	}
	transport, err = thrift.NewTBufferedTransportFactory(8192).GetTransport(transport)
	if err != nil {
		fmt.Printf("NewTransport failed. err: [%v]\n", err)
		return
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		fmt.Printf("Trasnport.Open failed. err: [%v]\n", err)
		return
	}

	protocolFactory := thrift.NewTCompactProtocolFactory()
	iprot := protocolFactory.GetProtocol(transport)
	oprot := protocolFactory.GetProtocol(transport)
	client := echo.NewEchoClient(thrift.NewTStandardClient(iprot, oprot))

	var res *echo.EchoRes
	res, err = client.Echo(context.Background(), &echo.EchoReq{
		Msg: strings.Join(os.Args[1:], " "),
	})
	if err != nil {
		fmt.Printf("client echo failed. err: [%v]\n", err)
		return
	}
	fmt.Printf("message from server: %v\n", res.GetMsg())
}
