package service

import (
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"path/filepath"
	"os"
)

func (s *Service) DeleteVM(ctx context.Context, req *pb.IDParcel) (*pb.Void, error) {
	vmPath := filepath.Join(s.config.VMDir, req.Value)
	s.runc.DeleteContainer(req.Value)
	os.RemoveAll(vmPath)
	return &pb.Void{}, nil
}
