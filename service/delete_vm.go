package service

import (
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"os"
	"path/filepath"
)

func (s *Service) DeleteVM(ctx context.Context, req *pb.TextParcel) (*pb.Void, error) {
	vmPath := filepath.Join(s.config.VMDir, req.Value)
	s.runc.DeleteContainer(req.Value)
	os.RemoveAll(vmPath)
	return &pb.Void{}, nil
}
