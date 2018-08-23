package service

import (
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"os"
	"path/filepath"
	"github.com/aemengo/bosh-runc-cpi/utils"
)

func (s *Service) DeleteVM(ctx context.Context, req *pb.TextParcel) (*pb.Void, error) {
	var (
		vmPath      = filepath.Join(s.config.VMDir, req.Value)
		rootFsPath  = filepath.Join(vmPath, "rootfs")
	)

	s.runc.DeleteContainer(req.Value)
	s.tearDownNetworking(vmPath)
	utils.RunCommand("umount", rootFsPath)
	os.RemoveAll(vmPath)
	return &pb.Void{}, nil
}
