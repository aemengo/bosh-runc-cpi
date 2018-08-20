package service

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"github.com/aemengo/bosh-runc-cpi/utils"
	"path/filepath"
	"sync"
	"strings"
)

func (s *Service) CheckpointVM(ctx context.Context, req *pb.VMFilterOpts) (*pb.Void, error) {
	vmIDs, err := s.vmIDs(req)
	if err != nil {
		return nil, err
	}

	wg := &sync.WaitGroup{}
	errChan := make(chan error, 1)

	for _, vmID := range vmIDs {
		if ok, _ := s.runc.HasContainer(vmID); !ok {
			continue
		}

		wg.Add(1)
		go s.checkpointVM(wg, vmID, errChan)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	errors := []string{}
	for err := range errChan {
		errors = append(errors, err.Error())
	}

	if len(errors) != 0 {
		return nil, fmt.Errorf("failed to perform checkpoint: %s", strings.Join(errors, "\n"))
	}

	return &pb.Void{}, nil
}

func (s *Service) checkpointVM(wg *sync.WaitGroup, vmID string, errChan chan error) {
	var (
		vmPath     = filepath.Join(s.config.VMDir, vmID)
		imagePath  = filepath.Join(vmPath, "criu", "image")
		workPath   = filepath.Join(vmPath, "criu", "work")
		parentPath = filepath.Join(vmPath, "criu", "parent")
		rootFsPath = filepath.Join(vmPath, "rootfs")
	)

	defer wg.Done()

	err := utils.MkdirAll(imagePath, workPath, parentPath)
	if err != nil {
		errChan <- fmt.Errorf("failed to checkpoint vm: %s: %s", vmID, err)
		return
	}

	// monit processes must be stopped before checkpoint
	// because the criu does not support processes made
	// outside the umbrella of the main container pid
	err = s.runc.StopProcesses(vmID)
	if err != nil {
		errChan <- fmt.Errorf("failed to checkpoint vm: %s: %s", vmID, err)
		return
	}

	// networking does not persist after restore anyway
	// so it is better to tear it down beforehand
	err = s.tearDownNetworking(vmPath)
	if err != nil {
		errChan <- fmt.Errorf("failed to checkpoint vm: %s: %s", vmID, err)
		return
	}

	err = s.runc.Checkpoint(vmID, imagePath, workPath, parentPath)
	if err != nil {
		errChan <- fmt.Errorf("failed to checkpoint vm: %s: %s", vmID, err)
		return
	}

	err = utils.RunCommand("umount", rootFsPath)
	if err != nil {
		errChan <- fmt.Errorf("failed to checkpoint vm: %s: %s", vmID, err)
		return
	}
}