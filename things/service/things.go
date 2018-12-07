package service

import (
	"context"
	"demo/pb"
	"errors"
)

type ThingsService struct {
}

func (*ThingsService) MethodY(ctx context.Context, req *pb.OK) (*pb.OK, error) {
	// logical, but return error or errs.New(code, msg)
	return nil, errors.New("woops")
}
