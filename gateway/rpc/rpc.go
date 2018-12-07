package rpc

import (
	"demo/pb"

	"github.com/arcplus/go-lib/os"
	"google.golang.org/grpc"
)

var UserClient = func() pb.UserClient {
	conn, err := grpc.Dial(os.Getenv("rpc_user_addr", "127.0.0.1:10000"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewUserClient(conn)
}()

var ThingsClient = func() pb.ThingsClient {
	conn, err := grpc.Dial(os.Getenv("rpc_things_addr", "127.0.0.1:10001"), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewThingsClient(conn)
}()
