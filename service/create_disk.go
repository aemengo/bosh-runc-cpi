package service

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"github.com/satori/go.uuid"
	"os"
	"path/filepath"
)

func (s *Service) CreateDisk(ctx context.Context, req *pb.NumberParcel) (*pb.TextParcel, error) {
	id := uuid.NewV4().String()
	diskPath := filepath.Join(s.config.DiskDir, id)

	err := os.MkdirAll(diskPath, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("disk directory could not be created: %s", err)
	}

	return &pb.TextParcel{Value: id}, nil
}
