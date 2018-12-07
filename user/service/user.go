package service

import (
	"context"

	"demo/pb"
)

type UserService struct {
}

func (*UserService) MethodX(ctx context.Context, req *pb.OK) (*pb.OK, error) {
	// logical
	return &pb.OK{}, nil
}
