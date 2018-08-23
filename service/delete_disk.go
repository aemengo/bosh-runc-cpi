package service

import (
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"os"
	"path/filepath"
)

func (s *Service) DeleteDisk(ctx context.Context, req *pb.TextParcel) (*pb.Void, error) {
	diskPath := filepath.Join(s.config.DiskDir, req.Value)

	os.RemoveAll(diskPath)
	return &pb.Void{}, nil
}