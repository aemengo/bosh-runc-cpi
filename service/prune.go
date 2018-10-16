package service

import (
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"github.com/aemengo/bosh-runc-cpi/utils"
)

func (s *Service) Prune(ctx context.Context, req *pb.Void) (*pb.Void, error) {
	return &pb.Void{}, utils.RunCommand("fstrim", s.config.WorkDir)
}