package service

import (
	"github.com/aemengo/bosh-runc-cpi/pb"
	"context"
)

func (s *Service) Ping(ctx context.Context, req *pb.Void) (*pb.IDParcel, error) {
	return &pb.IDParcel{Value: "pong"}, nil
}
