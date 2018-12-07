package main

import (
	"demo/pb"
	"demo/user/service"

	"github.com/arcplus/go-lib/micro"
	"github.com/arcplus/go-lib/mysql"
	"github.com/arcplus/go-lib/scaffold"
)

func main() {
	m := micro.New("user")

	m.AddResCloseFunc(mysql.Close)

	rpcServer := scaffold.NewGRPCServer()
	pb.RegisterUserServer(rpcServer, &service.UserService{})
	m.ServeGRPC(micro.Bind(":10000"), rpcServer)

	m.Start()
}
