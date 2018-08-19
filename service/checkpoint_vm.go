package service

import (
	"context"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"path/filepath"
	"github.com/aemengo/bosh-runc-cpi/utils"
)

func (s *Service) CheckpointVM(ctx context.Context, req *pb.VMFilterOpts) (*pb.Void, error) {
	vmIDs, err := s.vmIDs(req)
	if err != nil {
		return nil, err
	}

	for _, vmID := range vmIDs {
		var (
			vmPath = filepath.Join(s.config.VMDir, vmID)
			imagePath = filepath.Join(vmPath, "criu", "image")
			workPath = filepath.Join(vmPath, "criu", "work")
			parentPath = filepath.Join(vmPath, "criu", "parent")
		)

		err = utils.MkdirAll(imagePath, workPath, parentPath)
		if err != nil {
			return nil, err
		}

		err = s.runc.StopProcesses(vmID)
		if err != nil {
			return nil, err
		}

		err = s.tearDownNetworking(vmPath)
		if err != nil {
			return nil, err
		}

		err = s.runc.Checkpoint(vmID, imagePath, workPath, parentPath)
		if err != nil {
			return nil, err
		}
	}

	return &pb.Void{}, nil
}
