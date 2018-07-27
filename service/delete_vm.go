package service

import (
	"context"
	"github.com/aemengo/bosh-containerd-cpi/pb"
	"github.com/aemengo/bosh-containerd-cpi/utils"
	"os"
	"path/filepath"
)

func (s *Service) DeleteVM(ctx context.Context, req *pb.IDParcel) (*pb.Void, error) {
	var (
		vmPath     = filepath.Join(s.config.VMDir, req.Value)
		rootFsPath = filepath.Join(vmPath, "rootfs")
	)

	s.runc.DeleteContainer(req.Value)
	utils.RunCommand("umount", rootFsPath)
	os.RemoveAll(vmPath)
	return &pb.Void{}, nil
}
