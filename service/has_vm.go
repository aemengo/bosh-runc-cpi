package service

import (
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
)

func (s *Service) HasVM(ctx context.Context, req *pb.TextParcel) (*pb.TruthParcel, error) {
	exists, err := s.runc.HasContainer(req.Value)
	if err != nil {
		return nil, err
	}

	return &pb.TruthParcel{Value: exists}, nil
}