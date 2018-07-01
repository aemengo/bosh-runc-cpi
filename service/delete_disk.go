package service

import (
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"os"
	"path/filepath"
)

func (s *Service) DeleteDisk(ctx context.Context, req *pb.IDParcel) (*pb.Void, error) {
	diskPath := filepath.Join(s.config.DiskDir, req.Value)

	err := os.RemoveAll(diskPath)

	if err != nil {
		return nil, err
	} else {
		return &pb.Void{}, nil
	}
}