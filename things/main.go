package main

import (
	"demo/pb"
	"demo/things/service"

	"github.com/arcplus/go-lib/micro"
	"github.com/arcplus/go-lib/mysql"
	"github.com/arcplus/go-lib/scaffold"
)

func main() {
	m := micro.New("things")

	m.AddResCloseFunc(mysql.Close)

	rpcServer := scaffold.NewGRPCServer()
	pb.RegisterThingsServer(rpcServer, &service.ThingsService{})
	m.ServeGRPC(micro.Bind(":10001"), rpcServer)

	m.Start()
}
