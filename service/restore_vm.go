package service

import (
	"context"
	"fmt"
	"github.com/aemengo/bosh-runc-cpi/pb"
	"github.com/aemengo/bosh-runc-cpi/utils"
	"io/ioutil"
	"path/filepath"
	"sync"
	"strings"
)

func (s *Service) RestoreVM(ctx context.Context, req *pb.VMFilterOpts) (*pb.Void, error) {
	vmIDs, err := s.vmIDs(req)
	if err != nil {
		return nil, err
	}

	wg := &sync.WaitGroup{}
	errChan := make(chan error, 1)

	for _, vmID := range vmIDs {
		if ok, _ := s.runc.HasContainer(vmID); ok {
			continue
		}

		wg.Add(1)
		go s.restoreVM(wg, vmID, errChan)
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
		return nil, fmt.Errorf("failed to perform restore: %s", strings.Join(errors, "\n"))
	}

	return &pb.Void{}, nil
}

func (s *Service) restoreVM(wg *sync.WaitGroup, vmID string, errChan chan error) {
	var (
		vmPath            = filepath.Join(s.config.VMDir, vmID)
		imagePath         = filepath.Join(vmPath, "criu", "image")
		workPath          = filepath.Join(vmPath, "criu", "work")
		agentSettingsPath = filepath.Join(vmPath, "warden-cpi-agent-env.json")
		rootFsPath        = filepath.Join(vmPath, "rootfs")
		workDirPath       = filepath.Join(vmPath, "workdir")
		upperDirPath      = filepath.Join(vmPath, "upperdir")
		stemcellIDPath    = filepath.Join(vmPath, "stemcell-id")
		pidPath           = filepath.Join(vmPath, "pid")
	)

	defer wg.Done()

	contents, err := ioutil.ReadFile(stemcellIDPath)
	if err != nil {
		errChan <- fmt.Errorf("failed to read stemcell ID at: %s: %s", stemcellIDPath, err)
		return
	}

	stemcellPath := filepath.Join(s.config.StemcellDir, string(contents))

	err = utils.RunCommand("mount",
		"-t", "overlay",
		"-o", fmt.Sprintf("lowerdir=%s,upperdir=%s,workdir=%s", stemcellPath, upperDirPath, workDirPath),
		"overlay",
		rootFsPath,
	)
	if err != nil {
		errChan <- fmt.Errorf("failed to make rootfs: %s", err)
		return
	}

	err = s.runc.Restore(vmID, vmPath, imagePath, workPath, pidPath)
	if err != nil {
		errChan <- err
		return
	}

	agentSettings, err := ioutil.ReadFile(agentSettingsPath)
	if err != nil {
		errChan <- fmt.Errorf("failed to read agent settings file: %s", err)
		return
	}

	ip, mask, gatewayIP, err := extractNetValues(agentSettings)
	if err != nil {
		errChan <- fmt.Errorf("failed to extract network values: %s", err)
		return
	}

	err = s.configureNetworking(vmPath, pidPath, ip, mask, gatewayIP)
	if err != nil {
		errChan <- fmt.Errorf("failed to configure networker: %s", err)
		return
	}

	err = s.runc.StartProcesses(vmID)
	if err != nil {
		errChan <- err
		return
	}
}
