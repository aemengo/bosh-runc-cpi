package service

import (
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"path/filepath"
	"github.com/aemengo/bosh-runc-cpi/utils"
)

func (s *Service) HasDisk(ctx context.Context, req *pb.IDParcel) (*pb.TruthParcel, error) {
	diskPath := filepath.Join(s.config.DiskDir, req.Value)
	exists := utils.Exists(diskPath)
	return &pb.TruthParcel{Value: exists}, nil
}