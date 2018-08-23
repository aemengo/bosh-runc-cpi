package service

import (
	"github.com/aemengo/bosh-runc-cpi/pb"
	"context"
)

func (s *Service) Ping(ctx context.Context, req *pb.Void) (*pb.TextParcel, error) {
	return &pb.TextParcel{Value: "pong"}, nil
}
